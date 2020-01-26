package nixbuild

import (
	"embly/pkg/config"
	"embly/pkg/embly"
	"embly/pkg/filesystem"
	"embly/pkg/lucet"
	_ "embly/pkg/nixbuild/statik"
	"fmt"
	"io/ioutil"

	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/rakyll/statik/fs"
)

type Builder struct {
	emblyDir string
}

func NewBuilder() (b *Builder, err error) {
	b = &Builder{}
	if err = embly.CreateHomeDir(); err != nil {
		return
	}
	b.emblyDir, err = embly.EmblyDir()
	if err != nil {
		return
	}
	return
}

func (b *Builder) EmblyLoc(path string) string {
	return filepath.Join(b.emblyDir, path)
}

func (b *Builder) CleanAllDependencies() (err error) {
	dirs, err := filepath.Glob(b.EmblyLoc("./nix/") + "*.*")
	if err != nil {
		return
	}
	for _, file := range dirs {
		if err = os.RemoveAll(file); err != nil {
			return
		}
	}
	return
}

func (b *Builder) BuildRust() (err error) {
	defer func() {
		err = errors.WithStack(err)
	}()
	if err = b.WriteNixFiles(); err != nil {
		return
	}

	cfg, err := config.New("../../examples/kv")
	if err != nil {
		return
	}

	project := filesystem.NewProject(cfg)
	dir, err := project.CopyFunctionSourcesToBuild(b.EmblyLoc("./build"), "main")
	if err != nil {
		return
	}
	defer os.RemoveAll(dir)
	result, err := ioutil.TempDir(b.EmblyLoc("./result/"), "")
	if err != nil {
		return
	}
	{
		cmd := exec.Command(
			"nix", "run", "-i",
			"-f", b.EmblyLoc("./nix/rust-run.nix"),
			"--keep", "CARGO_TARGET_DIR",
			"--keep", "CARGO_HOME",
			"-c",
			"cargo", "build", "--release", "--target=wasm32-wasi",
			"-Z", "unstable-options", "--out-dir", result,
		)
		cmd.Dir = dir
		cmd.Env = []string{
			// this is needed for `nix run` it doesn't get passed into the command context
			fmt.Sprintf("NIX_PATH=%s", os.Getenv("NIX_PATH")),

			fmt.Sprintf("CARGO_TARGET_DIR=%s", b.EmblyLoc("./build_context/rust_target")),
			fmt.Sprintf("CARGO_HOME=%s", b.EmblyLoc("./build_context/cargo_home")),
		}
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if err = cmd.Run(); err != nil {
			return
		}
	}
	{
		var loc string
		loc, err = lucet.WriteBindingsFile()
		if err != nil {
			return
		}
		// --target x86_64-apple-darwin \
		cmd := exec.Command(
			"nix", "run", "-i",
			"-f", b.EmblyLoc("./nix/wasm-run.nix"),
			"--keep", "LDFLAGS",
			"--keep", "LD",
			"-c",
			"sh", "-c",
			// "env",
			`set -x && wasm-strip *.wasm \
			&& lucetc --bindings `+loc+`\
			--emit so \
			--opt-level 0 *.wasm \
			&& strip a.out`,
		)
		cmd.Dir = result
		cmd.Env = []string{
			// this is needed for `nix run` it doesn't get passed into the command context
			fmt.Sprintf("NIX_PATH=%s", os.Getenv("NIX_PATH")),
			// 	"LDFLAGS=-dylib -dead_strip -export_dynamic -undefined dynamic_lookup",
			// 	"LD=/nix/store/qlqy7maq14k4wxlc16yjikf12x5b0dln-cctools/bin/ld",
		}
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if err = cmd.Run(); err != nil {
			return
		}
	}

	return
}

func (b *Builder) WriteNixFiles() (err error) {
	defer func() {
		err = errors.WithStack(err)
	}()

	staticFiles, err := fs.New()
	if err != nil {
		return
	}
	if err = fs.Walk(
		staticFiles, "/",
		func(path string, fi os.FileInfo, e error) (err error) {
			defer func() {
				err = errors.WithStack(err)
			}()

			if e != nil || fi.IsDir() {
				return e
			}
			file, err := staticFiles.Open(path)
			if err != nil {
				return
			}
			to, err := os.OpenFile(b.EmblyLoc("./nix/"+fi.Name()),
				os.O_TRUNC|os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				return
			}
			defer to.Close()
			if _, err = io.Copy(to, file); err != nil {
				return
			}
			return
		},
	); err != nil {
		return
	}
	return
}

func (b *Builder) DownloadDependencies(items []string) (err error) {
	for _, item := range items {
		if _, err = os.Stat(b.EmblyLoc("./nix/" + item)); err == nil {
			// directory exists, continue to other dependencies
			continue
		}
		if err = b.BuildFile(item); err != nil {
			return
		}
	}
	return
}

func (b *Builder) BuildFile(name string) (err error) {
	defer func() {
		err = errors.WithStack(err)
	}()

	staticFiles, err := fs.New()
	if err != nil {
		return
	}
	f, err := staticFiles.Open(fmt.Sprintf("/%s.nix", name))
	if err != nil {
		return
	}

	cmd := exec.Command("nix-build", "-", "-o", name)
	emblyDir, err := embly.EmblyDir()
	if err != nil {
		return
	}
	cmd.Dir = filepath.Join(emblyDir, "./nix")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	if _, err = io.Copy(stdin, f); err != nil {
		return
	}
	stdin.Close()

	if err = cmd.Start(); err != nil {
		return
	}
	if err = cmd.Wait(); err != nil {
		return
	}

	return nil

}
