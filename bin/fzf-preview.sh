#!/usr/bin/env bash
#
# The purpose of this script is to demonstrate how to preview a file or an
# image in the preview window of fzf.

file=$1
type=$(file --mime-type "$file")

if [[ ! $type =~ image/ ]]; then
  # Sometimes bat is installed as batcat.
  if command -v batcat > /dev/null; then
    batname="batcat"
  elif command -v bat > /dev/null; then
    batname="bat"
  else
    cat "$1"
    exit
  fi

  ${batname} --style="${BAT_STYLE:-numbers}" --color=always --pager=never -- "$file"
elif [[ $KITTY_WINDOW_ID ]]; then
  # 'memory' is the fastest option but if you want the image to be scrollable,
  # you have to use 'stream'
  kitty icat --clear --transfer-mode=memory --stdin=no --place="${FZF_PREVIEW_COLUMNS}x${FZF_PREVIEW_LINES}@0x0" "$file" | sed \$d
  echo -en "\e[m"
elif [[ -n $FZF_PREVIEW_PIXEL_WIDTH ]]; then
  convert "$file" -resize "${FZF_PREVIEW_PIXEL_WIDTH}x${FZF_PREVIEW_PIXEL_HEIGHT}>" -dither FloydSteinberg sixel:-
else
  file "$file"
fi
