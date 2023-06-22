import * as cryptography from "/static/js/lib/cryptography.js"
import * as cookie from "/static/js/lib/cookie.js"
import * as api from "/static/js/lib/api.js"
import * as helpers from "/static/js/lib/helpers.js"

// forms
async function logoutFormSubmit(event) {
    event.preventDefault()

    cookie.clearTokenCookie()

    window.location = "/login"
}

async function aliasFormSubmit(event, form) {
    event.preventDefault()
    const formData = helpers.processForm(form)

    const [domainId, domain] = formData.domainInfo.split("//")
    const randomCode = cryptography.randomHex(8)
    const alias = formData.serviceName + "." + randomCode

    const error = await api.newAlias(domainId, alias, formData.iconUrl)
    if (error !== null) return console.log(error)

    await navigator.clipboard.writeText(alias + "@" + domain)
    form.reset()
    window.location.reload()
}

async function existingAliasFormSubmit(event, form) {
    event.preventDefault()
    const formData = helpers.processForm(form)

    const [domainId, _] = formData.domainInfo.split("//")

    const error = await api.newAlias(domainId, formData.alias, formData.iconUrl)
    if (error !== null) return console.log(error)

    form.reset()
    window.location.reload()
}

async function domainFormSubmit(event, form) {
    event.preventDefault()
    const formData = helpers.processForm(form)

    const error = await api.newDomain(formData.domain)
    if (error !== null) return console.log(error)

    form.reset()
    window.location.reload()
}

async function aliasDeleteFormSubmit(event, form) {
    event.preventDefault()
    const formData = helpers.processForm(form)

    const error = await api.deleteAlias(formData.aliasId)
    if (error !== null) return console.log(error)

    form.reset()
    window.location.reload()
}

async function domainDeleteFormSubmit(event, form) {
    event.preventDefault()
    const formData = helpers.processForm(form)

    const error = await api.deleteDomain(formData.domainId)
    if (error !== null) return console.log(error)

    form.reset()
    window.location.reload()
}

async function aliasDisableFormSubmit(event, form) {
    event.preventDefault()
    const formData = helpers.processForm(form)

    const error = await api.updateAliasDisabledStatus(formData.aliasId, true)
    if (error !== null) return console.log(error)

    form.reset()
    window.location.reload()
}

async function domainDisableFormSubmit(event, form) {
    event.preventDefault()
    const formData = helpers.processForm(form)

    const error = await api.updateDomainDisabledStatus(formData.domainId, true)
    if (error !== null) return console.log(error)

    form.reset()
    window.location.reload()
}

async function aliasEnableFormSubmit(event, form) {
    event.preventDefault()
    const formData = helpers.processForm(form)

    const error = await api.updateAliasDisabledStatus(formData.aliasId, false)
    if (error !== null) return console.log(error)

    form.reset()
    window.location.reload()
}

async function domainEnableFormSubmit(event, form) {
    event.preventDefault()
    const formData = helpers.processForm(form)

    const error = await api.updateDomainDisabledStatus(formData.domainId, false)
    if (error !== null) return console.log(error)

    form.reset()
    window.location.reload()
}

const logoutForm = document.getElementById("logoutForm")
const aliasForm = document.getElementById("aliasForm")
const existingAliasForm = document.getElementById("existingAliasForm")
const domainForm = document.getElementById("domainForm")
const aliasDisableForms = document.querySelectorAll("#aliasDisableForm")
const domainDisableForms = document.querySelectorAll("#domainDisableForm")
const aliasEnableForms = document.querySelectorAll("#aliasEnableForm")
const domainEnableForms = document.querySelectorAll("#domainEnableForm")
const aliasDeleteForms = document.querySelectorAll("#aliasDeleteForm")
const domainDeleteForms = document.querySelectorAll("#domainDeleteForm")

if (aliasForm !== null)
    aliasForm.addEventListener("submit", (event) => aliasFormSubmit(event, aliasForm))

if (existingAliasForm !== null)
    existingAliasForm.addEventListener("submit", (event) => existingAliasFormSubmit(event, existingAliasForm))

if (logoutForm !== null)
    logoutForm.addEventListener("submit", logoutFormSubmit)

if (domainForm !== null)
    domainForm.addEventListener("submit", (event) => domainFormSubmit(event, domainForm))

for (let form of aliasDisableForms)
    form.addEventListener("submit", (event) => aliasDisableFormSubmit(event, form))

for (let form of domainDisableForms)
    form.addEventListener("submit", (event) => domainDisableFormSubmit(event, form))

for (let form of aliasEnableForms)
    form.addEventListener("submit", (event) => aliasEnableFormSubmit(event, form))

for (let form of domainEnableForms)
    form.addEventListener("submit", (event) => domainEnableFormSubmit(event, form))

for (let form of aliasDeleteForms)
    form.addEventListener("submit", (event) => aliasDeleteFormSubmit(event, form))

for (let form of domainDeleteForms)
    form.addEventListener("submit", (event) => domainDeleteFormSubmit(event, form))


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


// start cookie monitoring 
cookie.monitorTokenForExpiry()