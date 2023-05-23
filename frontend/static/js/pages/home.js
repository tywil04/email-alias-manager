import * as cryptography from "/static/js/lib/cryptography.js"
import * as cookie from "/static/js/lib/cookie.js"


// forms
async function logoutFormSubmit(event) {
    event.preventDefault()

    cookie.clearTokenCookie()

    window.location = "/login"
}

async function aliasFormSubmit(event, form) {
    event.preventDefault()

    const formData = new FormData(form)
    const serviceName = formData.get("serviceName")
    const iconUrl = formData.get("iconUrl")
    const domainIdAndDomain = formData.get("domainId")
    const domainParts = domainIdAndDomain.split("//")
    const domainId = domainParts[0]
    const domain = domainParts[1]

    const randomCode = cryptography.randomHex(8)
    const alias = serviceName + "." + randomCode

    const payload = JSON.stringify({domainId: domainId, alias: alias, iconUrl: iconUrl})

    const response = await fetch("/api/v1/user/aliases", {
        method: "POST",
        headers: {
            "Content-type": "application/json",
        },
        body: payload,
    })

    if (response.status === 200) {
        await navigator.clipboard.writeText(alias + "@" + domain)
        form.reset()
        window.location.reload()
    } else {
        console.log("Error!")
    }
}

async function existingAliasFormSubmit(event, form) {
    event.preventDefault()

    const formData = new FormData(form)
    const alias = formData.get("alias")

    const payload = JSON.stringify({alias: alias})

    const response = await fetch("/api/v1/user/aliases", {
        method: "POST",
        headers: {
            "Content-type": "application/json",
        },
        body: payload,
    })

    if (response.status === 200) {
        form.reset()
        window.location.reload()
    } else {
        console.log("Error!")
    }
}

async function domainFormSubmit(event, form) {
    event.preventDefault()

    const formData = new FormData(form)
    const domain = formData.get("domain")

    const payload = JSON.stringify({domain: domain})

    const response = await fetch("/api/v1/user/domains", {
        method: "POST",
        headers: {
            "Content-type": "application/json",
        },
        body: payload,
    })

    if (response.status === 200) {
        form.reset()
        window.location.reload()
    } else {
        console.log("Error!")
    }
}

async function aliasDisableFormSubmit(event, form) {
    event.preventDefault()

    const formData = new FormData(form)
    const aliasId = formData.get("aliasId")

    const payload = JSON.stringify({aliasId: aliasId, disabled: true})

    const response = await fetch("/api/v1/user/aliases", {
        method: "PATCH",
        headers: {
            "Content-type": "application/json",
        },
        body: payload,
    })

    if (response.status === 200) {
        form.reset()
        window.location.reload()
    } else {
        console.log("Error!")
    }
}

async function domainDisableFormSubmit(event, form) {
    event.preventDefault()

    const formData = new FormData(form)
    const domainId = formData.get("domainId")

    const payload = JSON.stringify({domainId: domainId, disabled: true})

    const response = await fetch("/api/v1/user/domains", {
        method: "PATCH",
        headers: {
            "Content-type": "application/json",
        },
        body: payload,
    })

    if (response.status === 200) {
        form.reset()
        window.location.reload()
    } else {
        console.log("Error!")
    }
}

async function aliasEnableFormSubmit(event, form) {
    event.preventDefault()

    const formData = new FormData(form)
    const aliasId = formData.get("aliasId")

    const payload = JSON.stringify({aliasId: aliasId, disabled: false})

    const response = await fetch("/api/v1/user/aliases", {
        method: "PATCH",
        headers: {
            "Content-type": "application/json",
        },
        body: payload,
    })

    if (response.status === 200) {
        form.reset()
        window.location.reload()
    } else {
        console.log("Error!")
    }
}

async function domainEnableFormSubmit(event, form) {
    event.preventDefault()

    const formData = new FormData(form)
    const domainId = formData.get("domainId")

    const payload = JSON.stringify({domainId: domainId, disabled: false})

    const response = await fetch("/api/v1/user/domains", {
        method: "PATCH",
        headers: {
            "Content-type": "application/json",
        },
        body: payload,
    })

    if (response.status === 200) {
        form.reset()
        window.location.reload()
    } else {
        console.log("Error!")
    }
}

const logoutForm = document.getElementById("logoutForm")
const aliasForm = document.getElementById("aliasForm")
const existingAliasForm = document.getElementById("existingAliasForm")
const domainForm = document.getElementById("domainForm")
const aliasDisableForms = document.querySelectorAll("#aliasDisableForm")
const domainDisableForms = document.querySelectorAll("#domainDisableForm")
const aliasEnableForms = document.querySelectorAll("#aliasEnableForm")
const domainEnableForms = document.querySelectorAll("#domainEnableForm")

logoutForm.addEventListener("submit", logoutFormSubmit)
aliasForm.addEventListener("submit", (event) => aliasFormSubmit(event, aliasForm))
existingAliasForm.addEventListener("submit", (event) => existingAliasFormSubmit(event, existingAliasForm))
domainForm.addEventListener("submit", (event) => domainFormSubmit(event, domainForm))

for (let form of aliasDisableForms) {
    form.addEventListener("submit", (event) => aliasDisableFormSubmit(event, form))
}

for (let form of domainDisableForms) {
    form.addEventListener("submit", (event) => domainDisableFormSubmit(event, form))
}

for (let form of aliasEnableForms) {
    form.addEventListener("submit", (event) => aliasEnableFormSubmit(event, form))
}

for (let form of domainEnableForms) {
    form.addEventListener("submit", (event) => domainEnableFormSubmit(event, form))
}

// other 
function togglePageParam(bool, flag) {
    let pageParams = []
    if (window.location.search.length > 0) {
        pageParams = window.location.search.replace("?", "").split("&")
    }

    if (bool) {
        pageParams.push(flag)
    } else {
        pageParams = pageParams.filter((value) => value !== flag)
    }

    if (pageParams.length > 0) {
        window.location = window.location.href.split("?")[0] + "?" + pageParams.join("&")
    } else {
        window.location.search = ""
        window.location.href = window.location.href.replace("?", "")
    }
}

const pageParams = window.location.search.replace("?", "").split("&")

const showDisabledDomainsParam = "showDisabledDomains"
const showDisabledAliasesParam = "showDisabledAliases"

const showDisabledDomainsFormToggle = document.getElementById("showDisabledDomainsFormToggle")
const showDisabledAliasesFormToggle = document.getElementById("showDisabledAliasesFormToggle")

if (pageParams.includes(showDisabledDomainsParam)) {
    showDisabledDomainsFormToggle.checked = true
}

if (pageParams.includes(showDisabledAliasesParam)) {
    showDisabledAliasesFormToggle.checked = true
}

showDisabledDomainsFormToggle.addEventListener("change", (event) => togglePageParam(event.target.checked, showDisabledDomainsParam))
showDisabledAliasesFormToggle.addEventListener("change", (event) => togglePageParam(event.target.checked, showDisabledAliasesParam))