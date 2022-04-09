:set number
:set relativenumber
:set autoindent
:set tabstop=2
:set shiftwidth=2
:set smarttab
:set softtabstop=2
:set mouse=a
:set vb
:set ruler


call plug#begin()

Plug 'https://github.com/preservim/nerdtree' " NerdTree
Plug 'https://github.com/tpope/vim-commentary' " For Commenting gcc & gc
Plug 'https://github.com/vim-airline/vim-airline' " Status bar
Plug 'https://github.com/ap/vim-css-color' " CSS Color Preview
Plug 'https://github.com/ryanoasis/vim-devicons' " Developer Icons
Plug 'https://github.com/preservim/tagbar' " Tagbar for code navigation
Plug 'https://github.com/terryma/vim-multiple-cursors' " CTRL + N for multiple cursors
Plug 'morhetz/gruvbox' " Color scheme
Plug 'neoclide/coc.nvim', {'branch': 'release'} " Conquer of Completion
Plug 'junegunn/fzf', { 'do': { -> fzf#install() } }
Plug 'mg979/vim-visual-multi', {'branch': 'master'} " multi cursor
Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' } " Go development plugin 

call plug#end()

:colorscheme gruvbox

source ~/.config/nvim/modules/coc.vim
source ~/.config/nvim/modules/go.vim

" No more Arrow Keys, deal with it
noremap <Up> <NOP>
noremap <Down> <NOP>
noremap <Left> <NOP>
noremap <Right> <NOP>

" NERD Tree {
nnoremap <C-f> :NERDTreeFocus<CR>
nnoremap <C-n> :NERDTree<CR>
nnoremap <C-t> :NERDTreeToggle<CR>

let g:NERDTreeDirArrowExpandable="+"
let g:NERDTreeDirArrowCollapsible="~"
" }

nmap <F8> :TagbarToggle<CR>

:set completeopt-=preview " For No Preview

" fzf {
" find file in directory
silent! nmap <C-p> :FZF<CR>
" }

" air-line {
let g:airline_powerline_fonts = 1

if !exists('g:airline_symbols')
    let g:airline_symbols = {}
endif

" airline symbols
let g:airline_left_sep = ''
let g:airline_left_alt_sep = ''
let g:airline_right_sep = ''
let g:airline_right_alt_sep = ''
let g:airline_symbols.branch = ''
let g:airline_symbols.readonly = ''
let g:airline_symbols.linenr = ''
" }

