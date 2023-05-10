const uploadTitle = document.getElementById('title');
const uploadSubtitle = document.getElementById('subtitle');
const uploadAuthorName = document.getElementById('authorName');
const uploadPublishDate = document.getElementById('date');
const uploadAuthorPhoto = document.querySelector(".author-photo-form__input");
const uploadTinyPostIMG = document.querySelector(".input-hero-image-tiny__input");
const uploadPostIMG = document.querySelector(".input-hero-image__input");

const removeAuthorPhoto = document.querySelector(".author-photo-form__remove-button");
const removeTinyPostIMG = document.querySelector(".remove-tiny-hero");
const removePostIMG = document.querySelector(".remove-hero");

const publishButton = document.querySelector(".publish-button");

uploadTitle.addEventListener(
  "change" , 
  () => {
    let title = document.getElementById('title').value;
    let defaultTitle = 'New Post';
    if (title !== '' && title.length < 25) {
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
  "change",  
  () => {
    let subtitle = document.getElementById('subtitle').value;
    let defaultSubtitle = 'Please, enter any description';
    if (subtitle !== '' && subtitle.length < 60){
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
  "change",   
  () => {
    let name = document.getElementById('authorName').value;
    let defaultName = 'Enter author name';
    if (name !== '' && name.length < 25){
      document.getElementById('authorNameVisualTiny').innerHTML = name;
    }    
    else {
      document.getElementById('authorNameVisualTiny').innerHTML = defaultName;
    }
  }
)

uploadPublishDate.addEventListener(
  "change", 
  () => {
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
  "change",  
  () => {
    const previewPostCardAuthorPhoto = document.querySelector(".post-card-info__photo");
    const previewInput = document.querySelector(".preview-author-photo");
    const file = document.querySelector(".author-photo-form__input").files[0];
    const reader = new FileReader();
      reader.addEventListener(
        "load",
        () => {
          previewPostCardAuthorPhoto.src = reader.result;
          previewInput.src = reader.result;
        },
        false
      );
    document.querySelector(".author-photo-form__remove-button").classList.add("remove-button__remove-button-show");
    uploadAuthorPhotoButton = document.getElementById("uploadAuthorPhotoButton");
    uploadAuthorPhotoButton.innerHTML = 'Upload New';
    uploadAuthorPhotoButton.classList.add("author-photo-form__upload-button-view");
    document.querySelector(".upload-button__icon").classList.add("upload-button__icon-view");
    if (file) {
      reader.readAsDataURL(file);
    }
  }
)

removeAuthorPhoto.addEventListener(
  "click",
  () => {
    document.querySelector(".author-photo-form__remove-button").classList.remove("remove-button__remove-button-show");
    uploadAuthorPhotoButton = document.getElementById("uploadAuthorPhotoButton");
    uploadAuthorPhotoButton.innerHTML = 'Upload';
    uploadAuthorPhotoButton.classList.remove("author-photo-form__upload-button-view");
    document.querySelector(".upload-button__icon").classList.remove("upload-button__icon-view");
    const previewPostCardAuthorPhoto = document.querySelector(".post-card-info__photo");
    const previewInput = document.querySelector(".preview-author-photo");
    defaultAuthorPhoto = '/static/img/author0.png';
    previewPostCardAuthorPhoto.src = defaultAuthorPhoto;
    previewInput.src = defaultAuthorPhoto;

  }
)

uploadTinyPostIMG.addEventListener(
  "change",
  () => {
    const previewPostCardPhoto = document.querySelector(".post-card__photo");
    const previewInput = document.querySelector(".upload-place-tiny__img");
    const file = document.querySelector(".input-hero-image-tiny__input").files[0];
    const reader = new FileReader();
    reader.addEventListener(
      "load",
      () => {
        previewPostCardPhoto.src = reader.result;
        previewInput.src = reader.result;
        document.querySelector(".tiny-img-buttons").classList.add("tiny-img-buttons-show");
        document.querySelector(".input-hero-image-tiny__sign").classList.add("input-hero-image-tiny__sign-remove");
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
    const defaultPostTinyIMG = "/static/img/postphototiny0.png";
    const previewPostCardPhoto = document.querySelector(".post-card__photo");
    const previewInput = document.querySelector(".upload-place-tiny__img");
    previewPostCardPhoto.src = defaultPostTinyIMG;
    previewInput.src = defaultPostTinyIMG;
    document.querySelector(".tiny-img-buttons").classList.remove("tiny-img-buttons-show");
    document.querySelector(".input-hero-image-tiny__sign").classList.remove("input-hero-image-tiny__sign-remove");
  }
)

uploadPostIMG.addEventListener(
  "change",
  () => {
    const previewPostCardPhoto = document.querySelector(".article-preview-post-visual__photo");
    const previewInput = document.querySelector(".upload-place__img");
    const file = document.querySelector(".input-hero-image__input").files[0];
    const reader = new FileReader();
    reader.addEventListener(
      "load",
      () => {
        previewPostCardPhoto.src = reader.result;
        previewInput.src = reader.result;
        document.querySelector(".img-buttons").classList.add("img-buttons-show");
        document.querySelector(".input-hero-image__sign").classList.add("input-hero-image__sign-remove");
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
    const defaultPostIMG = "/static/img/postphoto0.png";
    const previewPostCardPhoto = document.querySelector(".article-preview-post-visual__photo");
    const previewInput = document.querySelector(".upload-place__img");
    document.querySelector(".img-buttons").classList.remove("img-buttons-show");
    document.querySelector(".input-hero-image__sign").classList.remove("input-hero-image__sign-remove");
    previewPostCardPhoto.src = defaultPostIMG;
    previewInput.src = defaultPostIMG;
  }
)

publishButton.addEventListener(
  "click",
  async () => {
    const response = await fetch(window.location.href);
    if (response.ok) {
      const reader = new FileReader();
      let titleInput = document.getElementById('title').value;
      let subtitleInput = document.getElementById('subtitle').value;
      let postImgInput = document.querySelector(".input-hero-image-tiny__input").files[0].name;
      let authorNameInput = document.querySelector(".author-photo-form__input").files[0].name;
      let authorIMGInput = document.getElementById('authorName').value;
      let previewIMGInput = document.querySelector(".input-hero-image__input").files[0].name
      let publishDateInput = document.querySelector(".input-publish-date__note").value;

      reader.addEventListener(
        "load",
        () => {
            
        },
        false
      );

      let data = {
        title: titleInput,
        subtitle: subtitleInput,
        postImg: postImgInput, 
        authorName: authorNameInput, 
        authorIMG: authorIMGInput,
        previewIMG: previewIMGInput,
        publishDate: publishDateInput,
      }

      let json = JSON.stringify(data);
      console.log(json);
    } else {
      console.log("Ошибка HTTP" + response);
    }    
  }
)
