# `cscope_lsp`

This LSP implementation uses cscope to get results quickly.

## Installation

```shell
go install github.com/dhananjaylatkar/cscope_lsp@latest
```

## Neovim config

```lua
local function start_cscope_lsp()
  local root_files =
    { "cscope.out", "cscope.files", "cscope.in.out", "cscope.out.in", "cscope.out.po", "cscope.po.out" }
  local paths = vim.fs.find(root_files, { stop = vim.env.HOME })
  local root_dir = vim.fs.dirname(paths[1])

  if root_dir then
    vim.lsp.start({
      name = "cscope_lsp",
      cmd = { "cscope_lsp" },
      root_dir = root_dir,
      filetypes = { "c", "h", "cpp", "hpp" },
    })
  end
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = { "c", "h", "cpp", "hpp" },
  desc = "Start cscope_lsp",
  callback = start_cscope_lsp,
})
```

## Requirements

1. `cscope` is installed.
2. `cscope.out` is created and updated.

## Supported Capabilities

1. textDocument/definition
2. textDocument/references

## Thanks

Used [educationalsp](https://github.com/tjdevries/educationalsp) as starter template.
