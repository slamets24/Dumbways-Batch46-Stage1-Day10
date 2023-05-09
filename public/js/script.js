function getData() {
    let name = document.getElementById("name").value
    let email = document.getElementById("email").value
    let phone = document.getElementById("phone").value
    let subject = document.getElementById("subject").value
    let message = document.getElementById("message").value


if(name == ""){
    return alert("Nama Harus Diisi")
}else if(email == ""){
    return alert("Email Harus Diisi")
}else if(phone == ""){
    return alert("phone Harus Diisi")
}else if(subject == ""){
    return alert("Subject Harus Diisi")
} else if (message == ""){
    return alert("Message Harus Diisi")
}

const destination = "slametseptiawan6@gmail.com"
let a = document.createElement("a")
a.href = `mailto:${destination}?subject=${subject}&body= Hallo nama saya ${name} , saya ingin ${message}, bisakah anda menghubungi saya di ${phone}`
a.click()

let data = {
    nama: name,
    email: email,
    telp: phone,
    subject: subject,
    pesan: message,
}

console.log(data)

}


