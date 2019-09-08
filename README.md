# Tao of Power
# Using R.L. Wing translation

This will randomly pick and present you with a passage from the Tao of Power
Note: There are 81 passages
I'm just fiddling and learning with Google Functions and Golang

* Original OCR
```
ocrmypdf --force-ocr TaoOfPower.pdf TaoOfPower.ocr.pdf
pdftotext TaoOfPower.ocr.pdf
```

* Google Functions Notes
```
gcloud functions deploy Tao --runtime go111 --trigger-http
gcloud functions list
gcloud functions describe Tao
```

* Sources
    * https://cloud.google.com/functions/
        * https://cloudonair.withgoogle.com/events/americas/watch?talk=cloud-functions
    * https://benjamincongdon.me/blog/2019/01/21/Getting-Started-with-Golang-Google-Cloud-Functions/
