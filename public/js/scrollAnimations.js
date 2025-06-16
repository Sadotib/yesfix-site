document.addEventListener('DOMContentLoaded', function () {
    const observerOptions = {
        threshold: 0.3,
    };

    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                const el = entry.target;
                const anim = el.getAttribute('data-animate');
                el.classList.remove('opacity-0');
                el.classList.add('animate-' + anim);
                observer.unobserve(el); // Run only once
            }
        });
    }, observerOptions);

    document.querySelectorAll('.js-scroll').forEach(el => {
        observer.observe(el);
    });
});
