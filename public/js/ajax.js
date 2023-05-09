const testimonial = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()

    xhr.open("GET", "https://api.npoint.io/e0f578cd66ad3bf87d1c", true)

    xhr.onload = function () {
        if (xhr.status == 200) {
            resolve(JSON.parse(xhr.response))
        } else {
            reject("Error loading data")
        }
    }

    xhr.onerror = function () {
        reject("Network Error")
    }

    xhr.send()
})

async function getAllTestimonial() {
    const response = await testimonial
    console.log(response)

    let testimonialForHtml = ""
    response.forEach((data) => {
        testimonialForHtml += `<div class="testimonial">
            <img src="${data.image}" class="profile-testimonial" />
            <p class="quote">"${data.comment}"</p>
            <p class="author">- ${data.name}</p>
            <p class="author">${data.rating} <i class="fa-solid fa-star"></i></p>
        </div>`
    })

    document.getElementById("testimonials").innerHTML = testimonialForHtml
}
getAllTestimonial()


async function filterTestimonials(rating) {
    const response = await testimonial
    let testimonialHTML = '';

    const testimonialFiltered = response.filter(function (data) {
        return data.rating === rating;
    })

    if (testimonialFiltered.length === 0) {
        testimonialHTML = `<h1> Data not found! </h1>`
    } else {
        testimonialFiltered.forEach(function (data) {
            testimonialHTML += `<div class="testimonial">
                <img src="${data.image}" class="profile-testimonial" />
                <p class="quote">"${data.comment}"</p>
                <p class="author">- ${data.name}</p>
                <p class="author">${data.rating} <i class="fa-solid fa-star"></i></p>
            </div>`
        })
    }

    document.getElementById('testimonials').innerHTML = testimonialHTML;
}