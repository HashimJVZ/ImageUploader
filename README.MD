# Image Uploader

## Description

This is a API based image uploader service. You can upload any images with the size limit of **5MB** and any format provided below:
`.png` `.jpg` `.jpeg` `.bmp` `.tiff` `.tif`

A **URL** to the image will be available on successful upload so that it can be embedded in web
pages or other web based documents.

The Image will be **deleted** automatically after **7 Days**.


## Installation

Use the docker image `image_uploader`to start the container and run the server.

You can rebuild the docker image by following steps:
1. Change directory into project directory.
2. Create docker Image using
```
docker build . -t <docker_image_tag>
```
replace `<docker_image_tag>` with any name.

You can also run the service locally by going to
```
http://localhost:8081
```
shows a webpage which will act as Interface.

You can change the **server_address** and **port** in the `.env` file.

## Working
1. Open the webpage `localhost:8081`.
2. Click on `Choose File` Button.
3. Select any Image you want to upload and click open.
4. You can click `upload` button to upload the image.
5. You can copy the `Image link` to preview the file if it uploaded successfully.

Example of an `Image link`
```
http://localhost:8081/images/2529839546.jpg
```

**File name** will be generated and will add the extension of the original file.

All the images will be saved under `static/images/` directory.

Logs will be printed to the default logger.