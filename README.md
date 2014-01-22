# Roka

Roka (concatenation of **Ro**maji and **Ka**na) is a very simple command line program to lookup Japanese kana using romaji.

## Install

  go get github.com/nwjlyons/roka

## Usage

~~~
  Usage: roka [options...] <romaji character>

  Options:
    -h Display Hiragana.
    -k Display Katakana.
~~~

This is what happens when you run roka:

~~~
  roka -h -k a
  Hiragana: あ
  Katakana: ア
~~~