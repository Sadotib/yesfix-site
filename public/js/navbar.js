document.addEventListener("DOMContentLoaded", function () {
  // Generic smooth scroll with offset (for sticky navbar)
  function scrollWithOffset(id, offset = 80) {
    const target = document.getElementById(id);
    if (!target) return;

    const y = target.getBoundingClientRect().top + window.pageYOffset - offset;
    window.scrollTo({ top: y, behavior: "smooth" });
  }

  // Pre-Book Now buttons
  document.querySelectorAll(".scroll-prebook").forEach(btn => {
    btn.addEventListener("click", (e) => {
      e.preventDefault();
      scrollWithOffset("prebook", 100); // Adjust offset as needed
    });
  });

  // Contact Us buttons
  document.querySelectorAll(".scroll-contact").forEach(btn => {
    btn.addEventListener("click", (e) => {
      e.preventDefault();
      scrollWithOffset("contact", 100);
    });
  });

  // Collapse mobile menu on link click
  document.querySelectorAll(".mobile-link").forEach(link => {
    link.addEventListener("click", () => {
      const menu = document.getElementById("mobile-menu");
      if (menu) menu.classList.add("hidden");
    });
  });

  // Toggle mobile menu open/close
  const menuToggle = document.getElementById("menu-toggle");
  const mobileMenu = document.getElementById("mobile-menu");
  if (menuToggle && mobileMenu) {
    menuToggle.addEventListener("click", () => {
      mobileMenu.classList.toggle("hidden");
    });
  }
});
