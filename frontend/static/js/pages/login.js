import * as cryptography from "/static/js/lib/cryptography.js"
import * as cookie from "/static/js/lib/cookie.js"
import * as helpers from "/static/js/lib/helpers.js"

async function formSubmit(event, form) {
    event.preventDefault()
    const formData = helpers.processForm(form)

    const token = await cryptography.generateToken(
        formData,username, 
        formData.password
    )
    
    cookie.setTokenCookie(token)

    form.reset()
    window.location = "/home"
}

let form = document.getElementById("form")

form.addEventListener("submit", (event) => formSubmit(event, form))