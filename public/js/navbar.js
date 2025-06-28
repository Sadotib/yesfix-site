// document.addEventListener("DOMContentLoaded", function () {
//   // Generic smooth scroll with offset (for sticky navbar)
//   function scrollWithOffset(id, offset = 80) {
//     const target = document.getElementById(id);
//     if (!target) return;

//     const y = target.getBoundingClientRect().top + window.pageYOffset - offset;
//     window.scrollTo({ top: y, behavior: "smooth" });
//   }

//   // Pre-Book Now buttons
//   document.querySelectorAll(".scroll-prebook").forEach(btn => {
//     btn.addEventListener("click", (e) => {
//       e.preventDefault();
//       scrollWithOffset("prebook", 100); // Adjust offset as needed
//     });
//   });

//   // Contact Us buttons
//   document.querySelectorAll(".scroll-contact").forEach(btn => {
//     btn.addEventListener("click", (e) => {
//       e.preventDefault();
//       scrollWithOffset("contact", 100);
//     });
//   });

//   // Collapse mobile menu on link click
//   document.querySelectorAll(".mobile-link").forEach(link => {
//     link.addEventListener("click", () => {
//       const menu = document.getElementById("mobile-menu");
//       if (menu) menu.classList.add("hidden");
//     });
//   });

//   // Toggle mobile menu open/close
//   const menuToggle = document.getElementById("menu-toggle");
//   const mobileMenu = document.getElementById("mobile-menu");
//   if (menuToggle && mobileMenu) {
//     menuToggle.addEventListener("click", () => {
//       mobileMenu.classList.toggle("hidden");
//     });
//   }
// });


// navbar.js

// Smooth scroll with offset
function scrollWithOffset(id, offset = 80) {
  const target = document.getElementById(id);
  if (!target) return;

  const y = target.getBoundingClientRect().top + window.pageYOffset - offset;
  window.scrollTo({ top: y, behavior: "smooth" });
}

// Toggle mobile menu
// function toggleMobileMenu() {
//   const mobileMenu = document.getElementById("mobile-menu");
//   if (mobileMenu) {
//     mobileMenu.classList.toggle("hidden");
//   }
// }

// Hide mobile menu (after clicking a link)
// function hideMobileMenu() {
//   const mobileMenu = document.getElementById("mobile-menu");
//   if (mobileMenu && !mobileMenu.classList.contains("hidden")) {
//     mobileMenu.classList.add("hidden");
//   }
// }

// Wait for DOM to be fully loaded
document.addEventListener("DOMContentLoaded", () => {
  // Attach click handlers for scroll buttons
  document.querySelectorAll(".scroll-prebook").forEach(btn => {
    btn.addEventListener("click", (e) => {
      e.preventDefault();
      scrollWithOffset("prebook", 100);
      hideMobileMenu();
    });
  });

  document.querySelectorAll(".scroll-contact").forEach(btn => {
    btn.addEventListener("click", (e) => {
      e.preventDefault();
      scrollWithOffset("contact", 100);
      hideMobileMenu();
    });
  });

  // Toggle menu on hamburger icon click
  // const menuToggle = document.getElementById("menu-toggle");
  // if (menuToggle) {
  //   menuToggle.addEventListener("click", toggleMobileMenu);
  // }

  // // Hide menu when clicking any mobile nav link
  // document.querySelectorAll(".mobile-link").forEach(link => {
  //   link.addEventListener("click", hideMobileMenu);
  // });
});
