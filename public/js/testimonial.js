const testimonialData = [{
    name: "Ahmad Sarutobi",
    comment: "Saya sangat puas dengan jasa yang saya beli dari perusahaan ini.",
    rating : 5,
    image: "https://images.unsplash.com/photo-1557862921-37829c790f19?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=871&q=80"
},{
    name: "Uzumaki Kula",
    comment: "Tim mereka sangat responsif dalam menanggapi pertanyaan dan permintaan saya, sehingga saya merasa sangat terbantu, KEREN.",
    rating: 4,
    image: "https://images.unsplash.com/photo-1564564321837-a57b7070ac4f?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=876&q=80"
},{
    name: "Habib Quirky",
    comment: "Saya sangat merekomendasikan perusahaan ini kepada siapa saja yang membutuhkan jasa yang mereka tawarkan.",
    rating: 4,
    image: "https://images.unsplash.com/photo-1600486913747-55e5470d6f40?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80"
},{
    name: "Fallenra Cendrawati",
    comment: "Saya tidak bisa berkata-kata cukup baik tentang pengalaman saya dengan jasa yang saya beli dari perusahaan ini. ",
    rating: 3,
    image: "https://images.unsplash.com/photo-1487412720507-e7ab37603c6f?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=871&q=80"
},{
    name: "Crooz Supardi",
    comment: "Saya pasti akan menggunakan jasa mereka lagi di masa depan dan merekomendasikan perusahaan ini kepada siapa saja yang membutuhkan jasa berkualitas tinggi.",
    rating: 5,
    image: "https://images.unsplash.com/photo-1548372290-8d01b6c8e78c?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=869&q=80"
}]

function showTestimonial() {
    let testimonialForHtml = ''

    testimonialData.forEach((item) => {
        testimonialForHtml += `<div class="testimonial">
        <img src=${item.image} class="profile-testimonial" />
        <p class="quote">${item.comment}</p>
        <p class="author">- ${item.name}</p>
        <p class="author">${item.rating}</p>
    </div>`
    })

    document.getElementById("testimonials").innerHTML = testimonialForHtml
}
showTestimonial()



function filterTestimonials(rating) {
    let testimonialForHtml = ''

    const dataFiltered = testimonialData.filter(function (data) {
        return data.rating === rating
    })
    console.log(dataFiltered)

    if (dataFiltered.length === 0) {
        testimonialForHtml = `<h3>Data not found ! </h3>`
    } else {
        dataFiltered.forEach((data) => {
            testimonialForHtml += `<div class="testimonial">
                <img src=${data.image} class="profile-testimonial" />
                <p class="quote">${data.comment}</p>
                <p class="author">- ${data.name}</p>
                <p class="author">${data.rating}</p>
            </div>`
        })
    }

    document.getElementById("testimonials").innerHTML = testimonialForHtml
}