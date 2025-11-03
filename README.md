# A-simple-archiver

A simple archiver for working on golang using video. With a variable-length code according to the algorithm

# Link

- **[Video](https://youtube.com/playlist?list=PLFAQFisfyqlX_pfTd09rT0zsl7BC6WTde&si=CW2Robm47IW1zvfz)** - example video

#  Installation

## Usage example
```bash
# an archiver of the type is being created Archive.exe
go build 

# specify that we are compressing the file example.txt with the flag method
./Archive pack -m vlc ./example.txt

# you can check if a vlc file has been created 
ls

# you can view the contents of the created file if you analyze it :)
cat example.vlc

# try to unpack the example.vlc, we will first delete the original
rm example.txt

./Archive unpack -m vlc ./example.vlc

# view the content. You will see "My name is Ted"
cat exmple.txt

# You can use another example in example.txt only without punctuation marks

