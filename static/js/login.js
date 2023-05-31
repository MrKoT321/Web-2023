const showPasswordButton = document.querySelector(".login-form-password__show-password");
const hidePasswordButton = document.querySelector(".login-form-password__hide-password");

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

