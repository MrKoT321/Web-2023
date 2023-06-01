const showPasswordButton = document.querySelector(".login-form-password__show-password");
const hidePasswordButton = document.querySelector(".login-form-password__hide-password");
const uploadEmail = document.querySelector(".login-form-email__input");
const uploadPassword = document.querySelector(".login-form-password__input");
const loginBuuton = document.querySelector(".login-form-button");
const invalidFormMessage = document.querySelector(".form-filled-incorrectly");
const emailError = document.querySelector(".login-form-email__error");
const passwordError = document.querySelector(".login-form-password__error");

uploadEmail.addEventListener(
  "input",
  () => {
    uploadEmail.style.borderBottom = "1px solid #2E2E2E";
    emailError.classList.add("hidden");
  }
)

uploadPassword.addEventListener(
  "input",
  () => {
    uploadPassword.style.borderBottom = "1px solid #2E2E2E";
    passwordError.classList.add("hidden");
  }
)

showPasswordButton.addEventListener(
  "click",
  () => {
    showPasswordButton.classList.add("login-form-password__show-password-replace");
    hidePasswordButton.classList.add("login-form-password__hide-password-replace");
    document.querySelector(".login-form-password__input").type = "text";
  }
)

hidePasswordButton.addEventListener(
    "click",
    () => {
      showPasswordButton.classList.remove("login-form-password__show-password-replace");
      hidePasswordButton.classList.remove("login-form-password__hide-password-replace");
      document.querySelector(".login-form-password__input").type = "password";
    }
  )

loginBuuton.addEventListener(
  "click",
  () => {
    emailError.classList.add("hidden");
    passwordError.classList.add("hidden");
    let textInputValid = "1px solid #2E2E2E";
    let textInputInvalid = "1px solid #E86961";
    invalidFormMessage.classList.add("form-filled-incorrectly-hidden");
    let formError = false;
    for (child of invalidFormMessage.children) {
      child.classList.add("hidden");
    }
    if (uploadEmail.value === "") {
      emailError.classList.remove("hidden");
      uploadEmail.style.borderBottom = textInputInvalid;
      formError = true;
    } else { uploadEmail.style.borderBottom = textInputValid }
    if (uploadPassword.value === "") {
      passwordError.classList.remove("hidden");
      uploadPassword.style.borderBottom = textInputInvalid;
      formError = true;
    } else { uploadPassword.style.borderBottom = textInputValid }
    if (formError) {
      invalidFormMessage.classList.remove("form-filled-incorrectly-hidden");
      for (child of invalidFormMessage.children) {
        child.classList.remove("hidden");
      }
    }
  }
)
