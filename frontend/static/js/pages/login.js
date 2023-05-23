import * as cryptography from "/static/js/lib/cryptography.js"
import * as cookie from "/static/js/lib/cookie.js"

async function formSubmit(event, form) {
    event.preventDefault()

    const formData = new FormData(form)
    const username = formData.get("username")
    const password = formData.get("password")

    const token = await cryptography.generateToken(username, password)
    
    cookie.setTokenCookie(token)

    form.reset()
    window.location = "/home"
}

let form = document.getElementById("form")

form.addEventListener("submit", (event) => formSubmit(event, form))