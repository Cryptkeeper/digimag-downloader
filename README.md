# digimag-downloader
A simple CLI tool for bulk downloading ["Digimag"](https://mydigimag.rrd.com/) publications. This allows you to download and archive magazine publications that would otherwise be online only, requiring a proprietary service for viewing.

Given a publication ID, `digimag-downloader` will download each page of _each issue_ as a PDF to its working directory using the file path `issue_(issue id)_((issue date))_page_(page #).pdf`.

The resulting (per page) PDF files can be merged using a [variety of tools](https://superuser.com/questions/54041/how-to-merge-pdfs-using-imagemagick-resolution-problem).

## Compiling
1. Clone the repository using git: `git clone https://github.com/Cryptkeeper/digimag-downloader`
2. Compile the project using go: `go build ./cmd/digimag-downloader`

## Usage
Once compiled, `digimag-downloader` requires a single argument: the publication ID.

`./digimag-downloader <publication ID>`

### Extracting Publication IDs
When loading the viewer application in your web browser, the publication ID is the value of the `m` parameter. This is a persistent value and should not change unless the publication is discontinued and restarted under a new name.

For example, `https://mydigimag.rrd.com/publication/?m=1&i=2&p=3`
* The publication ID (`m`) is 1
* The issue ID (`i`) is 2
* The current page index (`p`) is 3 (this parameter may be omitted assuming you're viewing the cover page)

## License
See [LICENSE](LICENSE).