document.addEventListener("DOMContentLoaded", () => {
    const likeButtons = document.querySelectorAll(".like-btn");

    likeButtons.forEach((btn) => {
        btn.addEventListener("click", () => {
            const icon = btn.querySelector(".like-icon");
            const text = btn.querySelector(".like-text");
            const isLiked = icon.classList.contains("fa-solid");

            if (isLiked) {
                icon.classList.remove("fa-solid", "text-red-500");
                icon.classList.add("fa-regular", "text-gray-400");
                text.textContent = "Like";
                text.classList.remove("text-red-500");
                text.classList.add("text-gray-600");
            } else {
                icon.classList.remove("fa-regular", "text-gray-400");
                icon.classList.add("fa-solid", "text-red-500");
                text.textContent = "Liked";
                text.classList.add("text-red-500");
                text.classList.remove("text-gray-600");

                confetti({
                    particleCount: 120,
                    spread: 70,
                    origin: { x: 0.5, y: 1, }

                });
            }
        });
    });

    const carousel = document.getElementById("videoSlides");
    const slides = carousel.querySelectorAll("div.min-w-full");
    const totalSlides = slides.length;
    let index = 0;

    const updateCarousel = () => {
        carousel.style.transform = `translateX(-${index * 100}%)`;

        slides.forEach((slide, i) => {
            if (i === index) {
                slide.classList.remove("blur-sm", "opacity-50", "scale-95");
                slide.classList.add("scale-100");
            } else {
                slide.classList.add("blur-sm", "opacity-50", "scale-95");
                slide.classList.remove("scale-100");
            }
        });
    };

    document.getElementById("prevBtn").addEventListener("click", () => {
        index = (index - 1 + totalSlides) % totalSlides;
        updateCarousel();
    });

    document.getElementById("nextBtn").addEventListener("click", () => {
        index = (index + 1) % totalSlides;
        updateCarousel();
    });

    // Initial state
    updateCarousel();

});