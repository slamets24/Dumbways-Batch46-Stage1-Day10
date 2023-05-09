class Testimonial {
    #quote = ""
    #image = ""

    constructor(quote, image) {
        this.#quote = quote
        this.#image = image
    }

    get quote() {
        return this.#quote
    }
    
    get image() {
        return this.#image
    }

    get forHtml() {
        return `<div class="testimonial">
                    <img src="${this.image}" />
                    <p class="quote">${this.quote}</p>
                    <p class="author">- ${this.author}</p>
                </div>`
    }
}

class AuthorTestimonial extends Testimonial {
    #author = ""

    constructor(author, quote, image){
        super(quote, image)
        this. #author = author
    }

    get author() {
        return this.#author
    }
}


const first_coment = new AuthorTestimonial("Ahmad Sarutobi", "Saya sangat puas dengan jasa yang saya beli dari perusahaan ini.", "https://images.unsplash.com/photo-1557862921-37829c790f19?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=871&q=80")

const second_coment = new AuthorTestimonial("Uzumaki Kula", "Tim mereka sangat responsif dalam menanggapi pertanyaan dan permintaan saya, sehingga saya merasa sangat terbantu, KEREN.", "https://images.unsplash.com/photo-1564564321837-a57b7070ac4f?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=876&q=80")

const third_coment = new AuthorTestimonial("Habib Quirky", "Saya sangat merekomendasikan perusahaan ini kepada siapa saja yang membutuhkan jasa yang mereka tawarkan.", "https://images.unsplash.com/photo-1600486913747-55e5470d6f40?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80")

const fourth_coment = new AuthorTestimonial("Fallenra Cendrawati", "Saya tidak bisa berkata-kata cukup baik tentang pengalaman saya dengan jasa yang saya beli dari perusahaan ini. ", "https://images.unsplash.com/photo-1487412720507-e7ab37603c6f?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=871&q=80")

const fifth_coment = new AuthorTestimonial("Crooz Supardi", "Saya pasti akan menggunakan jasa mereka lagi di masa depan dan merekomendasikan perusahaan ini kepada siapa saja yang membutuhkan jasa berkualitas tinggi.", "https://images.unsplash.com/photo-1548372290-8d01b6c8e78c?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=869&q=80")

let data = [first_coment, second_coment, third_coment, fourth_coment, fifth_coment]

let testimonialForHtml = ""

for (let i=0; i < data.length; i++){
    testimonialForHtml += data[i]. forHtml
}

document.getElementById("testimonials").innerHTML = testimonialForHtml
