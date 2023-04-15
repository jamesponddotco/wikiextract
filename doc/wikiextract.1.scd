wikiextract(1)

# NAME

wikiextract - A word extractor for Wikipedia articles

# SYNOPSIS

*wikiextract* [options...]

# DESCRIPTION

**wikiextract** is a word extractor for Wikipedia articles. It can extract
words bigger than 4 characters from a given Wikipedia page or list of pages and
save them to a file you can later use as the source for generating diceware
passwords.

# OPTIONS

*-u*, *--url*
	Provide the URL of the Wikipedia page from which you want to extract words.

*-l*, *--list-file*
	Provide a file containing a list of URLs of Wikipedia pages from which you want to extract words. The file should have one URL per line.

*-o*, *--output-file*
	Provide the output file path where the extracted words will be written.

*-h*, *--help*
	Show help message and quit.

# EXAMPLES

Extract words from a single Wikipedia page:

	wikiextract -u https://en.wikipedia.org/wiki/Example -o output.txt

Extract words from multiple Wikipedia pages listed in a file:

	wikiextract -l list_of_urls.txt -o output.txt

# AUTHORS

Maintained by James Pond <james@cipher.host>.
