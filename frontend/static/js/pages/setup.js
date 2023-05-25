import * as cryptography from "/static/js/lib/cryptography.js"
import * as cookie from "/static/js/lib/cookie.js"

async function formSubmit(event, form) {
    event.preventDefault()

    const formData = new FormData(form)
    const username = formData.get("username")
    const password = formData.get("password")
    const passwordConfirm = formData.get("passwordConfirm")

    if (password !== passwordConfirm) {
        console.log("Error!")
        return
    }

    const token = await cryptography.generateToken(username, password)

    const payload = JSON.stringify({token: token})

    const response = await fetch("/api/v1/user/setup", {
        method: "POST",
        headers: {
            "Content-type": "application/json",
        },
        body: payload,
    })
    
    if (response.status === 200) {
        cookie.setTokenCookie(token)
        form.reset()
        window.location = "/home"
    } else {
        console.log("Error!")
    }
}

let form = document.getElementById("form")

form.addEventListener("submit", (event) => formSubmit(event, form))