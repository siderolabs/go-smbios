# go-smbios

Some notes on using this package:

All possible smbios info is document in the spec [here](https://www.dmtf.org/sites/default/files/standards/documents/DSP0134_2.7.1.pdf).

The underlying DigitalOcean smbios library provides a struct for each smbios structure with the given format below.

```go
type Structure struct {
  Header    Header
  Formatted []byte
  Strings   []string
}
```

The `Formatted` section contains the entire byte slice of the structure minus the header, while the `Strings` section contains a list of all strings that are provided in the structure.
Note that strings are _**only**_ present if the `Value` type in the specification is of type "STRING".
But this means you can easily get string values by their index, so in the example table lower in the README one can assume that `Manufacturer` will be in `Strings[0]`.

The way the Structure struct works for a given type can also be kind of confusing.
For example, in the doc, the table for system info looks like the following:

<img src="docs/img/system info table.png" width="500" alt="system info table">

But it's important to note that if I'm after a field like `Wake-up Type`, I need to keep in mind that the `Formatted` byte slice is missing the first 4 bytes of the structure that are stripped out as header info.
So if `Wake-up Type`'s offset is 18h (which is the decimal value 24), I need to subtract 4 to get the correct offset location (which is decimal value 20).
Thus fetching `s.Formatted[20]` gives me the byte that points to the wake up value, and I can cross-check that with the info from the spec.
[This](https://www.prepressure.com/library/technology/ascii-binary-hex) site was also helpful if you don't have your hex->decimal translations memorized.
