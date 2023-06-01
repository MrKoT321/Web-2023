const uploadTitle = document.getElementById('title');
const uploadSubtitle = document.getElementById('subtitle');
const uploadAuthorName = document.getElementById('authorName');
const uploadPublishDate = document.getElementById('date');
const uploadAuthorPhoto = document.querySelector(".author-photo-form__input");
const uploadTinyPostIMG = document.querySelector(".input-hero-image-tiny__input");
const uploadPostIMG = document.querySelector(".input-hero-image__input");
const uploadContent = document.querySelector(".upload-post-article__input");

const removeAuthorPhoto = document.querySelector(".author-photo-form__remove-button");
const removeTinyPostIMG = document.querySelector(".remove-tiny-hero");
const removePostIMG = document.querySelector(".remove-hero");

const succsesFormMessage = document.querySelector(".form-filled-correctly");
const invalidFormMessage = document.querySelector(".form-filled-incorrectly");

const publishButton = document.querySelector(".publish-button");

const titleError = document.querySelector(".input-title__error");
const subtitleError = document.querySelector(".input-subtitle__error");
const authorNameError = document.querySelector(".input-author-name__error");
const authorPhotoError = document.querySelector(".input-author-photo__error");
const publishDateError = document.querySelector(".input-publish-date__error");
const postImgError = document.querySelector(".input-hero-image__error");
const previewImgError = document.querySelector(".input-hero-image-tiny__error");
const contentError = document.querySelector(".upload-post-article__error");

let postImgInput;
let authorIMGInput;
let previewIMGInput;

let authorIMGInputName;
let previewIMGInputName;
let postImgInputName;

uploadTitle.addEventListener(
  "input" , 
  () => {
    uploadTitle.style.borderBottom = "1px solid #2E2E2E";
    titleError.classList.add("hidden");
    let title = uploadTitle.value;
    let defaultTitle = 'New Post';
    if (title !== '') {
      document.getElementById('titleVisual').innerHTML = title;
      document.getElementById('titleVisualTiny').innerHTML = title;
    }
    else {
      document.getElementById('titleVisual').innerHTML = defaultTitle;
      document.getElementById('titleVisualTiny').innerHTML = defaultTitle;
    }
  }
)

uploadSubtitle.addEventListener(
  "input",  
  () => {
    uploadSubtitle.style.borderBottom = "1px solid #2E2E2E";
    subtitleError.classList.add("hidden");
    let subtitle = document.getElementById('subtitle').value;
    let defaultSubtitle = 'Please, enter any description';
    if (subtitle !== ''){
      document.getElementById('subtitleVisual').innerHTML = subtitle;
      document.getElementById('subtitleVisualTiny').innerHTML = subtitle;
    } 
    else {
      document.getElementById('subtitleVisual').innerHTML = defaultSubtitle;
      document.getElementById('subtitleVisualTiny').innerHTML = defaultSubtitle;
    }   
  }
)

uploadAuthorName.addEventListener(
  "input",   
  () => {
    uploadAuthorName.style.borderBottom = "1px solid #2E2E2E";
    authorNameError.classList.add("hidden");
    let name = document.getElementById('authorName').value;
    let defaultName = 'Enter author name';
    if (name !== ''){
      document.getElementById('authorNameVisualTiny').innerHTML = name;
    }    
    else {
      document.getElementById('authorNameVisualTiny').innerHTML = defaultName;
    }
  }
)

uploadPublishDate.addEventListener(
  "input", 
  () => {
    uploadPublishDate.style.borderBottom = "1px solid #2E2E2E";
    publishDateError.classList.add("hidden");
    let date = document.getElementById('date').value;
    let defaultDate = '01/05/2023';
    if (date !== ''){
      document.getElementById('dateVisuality').innerHTML = date;
    }    
    else {
      document.getElementById('dateVisuality').innerHTML = defaultDate;
    }
  }
)

uploadAuthorPhoto.addEventListener(
  "input",  
  () => {
    authorPhotoError.classList.add("hidden");
    const previewPostCardAuthorPhoto = document.querySelector(".post-card-info__photo");
    const previewInput = document.querySelector(".preview-author-photo");
    const file = document.querySelector(".author-photo-form__input").files[0];
    const reader = new FileReader();
      reader.addEventListener(
        "load",
        () => {
          previewPostCardAuthorPhoto.src = reader.result;
          previewInput.src = reader.result;
          authorIMGInput = reader.result.replace("data:", "").replace(/^.+,/, "");
          authorIMGInputName = file.name;
        },
        false
      );
    removeAuthorPhoto.classList.remove("hidden");
    uploadAuthorPhotoButton = document.getElementById("uploadAuthorPhotoButton");
    uploadAuthorPhotoButton.innerHTML = 'Upload New';
    uploadAuthorPhotoButton.classList.add("author-photo-form__upload-button-view");
    document.querySelector(".upload-button__icon").classList.remove("hidden");
    if (file) {
      reader.readAsDataURL(file);
    }
  }
)

removeAuthorPhoto.addEventListener(
  "click",
  () => {
    uploadAuthorPhoto.value = null;
    removeAuthorPhoto.classList.add("hidden");
    uploadAuthorPhotoButton = document.getElementById("uploadAuthorPhotoButton");
    uploadAuthorPhotoButton.innerHTML = 'Upload';
    uploadAuthorPhotoButton.classList.remove("author-photo-form__upload-button-view");
    document.querySelector(".upload-button__icon").classList.add("hidden");
    const previewPostCardAuthorPhoto = document.querySelector(".post-card-info__photo");
    const previewInput = document.querySelector(".preview-author-photo");
    previewPostCardAuthorPhoto.src = "/static/img/defaultpostphoto.png" ;
    previewInput.src = "/static/img/author0.png";

  }
)

uploadTinyPostIMG.addEventListener(
  "input",
  () => {
    previewImgError.classList.add("hidden");
    const previewPostCardPhoto = document.querySelector(".post-card__photo");
    const previewInput = document.querySelector(".upload-place-tiny__img");
    const file = document.querySelector(".input-hero-image-tiny__input").files[0];
    const reader = new FileReader();
    reader.addEventListener(
      "load",
      () => {
        previewPostCardPhoto.src = reader.result;
        previewInput.src = reader.result;
        previewIMGInput = reader.result.replace("data:", "").replace(/^.+,/, "");
        previewIMGInputName = file.name;
        document.querySelector(".tiny-img-buttons").classList.remove("hidden");
        document.querySelector(".input-hero-image-tiny__sign").classList.add("hidden");
      },
      false
    );

    if (file) {
      reader.readAsDataURL(file);
    }
  }
)

removeTinyPostIMG.addEventListener(
  "click",
  () => {
    uploadTinyPostIMG.value = null;
    const previewPostCardPhoto = document.querySelector(".post-card__photo");
    const previewInput = document.querySelector(".upload-place-tiny__img");
    previewPostCardPhoto.src = "/static/img/defaultpostphoto.png";
    previewInput.src = "/static/img/postphototiny0.png";
    document.querySelector(".tiny-img-buttons").classList.add("hidden");
    document.querySelector(".input-hero-image-tiny__sign").classList.remove("hidden");
  }
)

uploadPostIMG.addEventListener(
  "input",
  () => {
    postImgError.classList.add("hidden");
    const previewPostCardPhoto = document.querySelector(".article-preview-post-visual__photo");
    const previewInput = document.querySelector(".upload-place__img");
    const file = document.querySelector(".input-hero-image__input").files[0];
    const reader = new FileReader();
    reader.addEventListener(
      "load",
      () => {
        previewPostCardPhoto.src = reader.result;
        previewInput.src = reader.result;
        postImgInput = reader.result.replace("data:", "").replace(/^.+,/, "");
        postImgInputName = file.name;
        document.querySelector(".img-buttons").classList.remove("hidden");
        document.querySelector(".input-hero-image__sign").classList.add("hidden");
      },
      false
    );

    if (file) {
      reader.readAsDataURL(file);
    }
  }
)

removePostIMG.addEventListener(
  "click",
  () => {
    uploadPostIMG.value = null;
    const previewPostCardPhoto = document.querySelector(".article-preview-post-visual__photo");
    const previewInput = document.querySelector(".upload-place__img");
    document.querySelector(".img-buttons").classList.add("hidden");
    document.querySelector(".input-hero-image__sign").classList.remove("hidden");
    previewPostCardPhoto.src = "/static/img/defaultpostphoto.png";
    previewInput.src = "/static/img/postphoto0.png";
  }
)

uploadContent.addEventListener(
  "input",
  () => {
    contentError.classList.add("hidden");
    uploadContent.style.border = "1px solid #EAEAEA";
  }
)

function CheckValidInputs() {
  succsesFormMessage.classList.add("form-filled-correctly-hidden");
  invalidFormMessage.classList.add("form-filled-incorrectly-hidden");
  for (child of succsesFormMessage.children) {
    child.classList.add("hidden");
  } 
  for (child of invalidFormMessage.children) {
    child.classList.add("hidden");
  }
  let textInputValid = "1px solid #2E2E2E";
  let textInputInvalid = "1px solid #E86961";

  let formError = false;
  if (uploadTitle.value === "") {
    titleError.classList.remove("hidden");
    uploadTitle.style.borderBottom = textInputInvalid;
    formError = true;
  } else { uploadTitle.style.borderBottom = textInputValid }
  if (uploadSubtitle.value === "") {
    subtitleError.classList.remove("hidden");
    uploadSubtitle.style.borderBottom = textInputInvalid;
    formError = true;
  } else { uploadSubtitle.style.borderBottom = textInputValid }
  if (uploadAuthorName.value === "") {
    authorNameError.classList.remove("hidden");
    uploadAuthorName.style.borderBottom = textInputInvalid;
    formError = true;
  } else { uploadAuthorName.style.borderBottom = textInputValid }
  if (uploadPublishDate.value === "") {
    publishDateError.classList.remove("hidden");
    uploadPublishDate.style.borderBottom = textInputInvalid;
    formError = true;
  } else { uploadPublishDate.style.borderBottom = textInputValid }
  if (!uploadAuthorPhoto.files[0]) {
    authorPhotoError.classList.remove("hidden");
    formError = true;
  }
  if (!uploadPostIMG.files[0]) {
    postImgError.classList.remove("hidden");
    formError = true;
  }
  if (!uploadTinyPostIMG.files[0]) {
    previewImgError.classList.remove("hidden");
    formError = true;
  }
  if (uploadContent.value === "") {
    uploadContent.style.border = "2px solid #E86961";
    contentError.classList.remove("hidden");
    formError = true;
  } else { uploadContent.style.border = "1px solid #EAEAEA" }
  if (formError) {
    invalidFormMessage.classList.remove("form-filled-incorrectly-hidden");
    for (child of invalidFormMessage.children) {
      child.classList.remove("hidden");
    }
  } else {
    succsesFormMessage.classList.remove("form-filled-correctly-hidden");
    for (child of succsesFormMessage.children) {
      child.classList.remove("hidden");
    }
  }
  return formError
}

publishButton.addEventListener(
  "click",
  () => {
    if (!CheckValidInputs()){
      const data = {
        title: uploadTitle.value,
        subtitle: uploadSubtitle.value,
        postIMG: postImgInput, 
        postIMGName: postImgInputName, 
        authorName: uploadAuthorName.value, 
        authorIMG: authorIMGInput,
        authorIMGName: authorIMGInputName,
        previewIMG: previewIMGInput,
        previewIMGName: previewIMGInputName,
        publishDate: uploadPublishDate.value,
        content: document.querySelector(".upload-post-article__input").value,
      }
      console.log(JSON.stringify(data, null, "\t"));
      doFecth(data);
    }
    
  }    
)

async function doFecth(data) {
  const response = await fetch("/api/post", {
    method: "POST",
    headers: {
      "Content-Type": "application/json;charset=utf-8"
    },
    body: JSON.stringify(data, null, "\t")
  });
  if (!response.ok) {
    alert("Ошибка HTTP: " + response.status);
  }
}
