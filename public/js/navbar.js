// Smooth scroll with offset
function scrollWithOffset(id, offset = 250) {
  const target = document.getElementById(id);
  if (!target) return;

  const y = target.getBoundingClientRect().top + window.pageYOffset - offset;
  window.scrollTo({ top: y, behavior: "smooth" });
}

// Hide mobile menu (after clicking a link)
function hideMobileMenu() {
  const mobileMenu = document.getElementById("mobile-menu");
  if (mobileMenu && !mobileMenu.classList.contains("hidden")) {
    mobileMenu.classList.add("hidden");
    const menuIcon = document.getElementById("menu-icon");
    if (menuIcon) {
      menuIcon.textContent = "☰"; // Reset icon to hamburger
    }
  }
}

// Toggle mobile menu
function toggleMobileMenuIcon() {
  const mobileMenu = document.getElementById("mobile-menu");
  const menuIcon = document.getElementById("menu-icon");

  if (mobileMenu && menuIcon) {
    const isHidden = mobileMenu.classList.toggle("hidden");
    menuIcon.textContent = isHidden ? "☰" : "✕";
  }
}

if (
  menu &&
  !menu.classList.contains("hidden") &&
  !menu.contains(e.target) &&
  !toggleBtn.contains(e.target)
) {
  menu.classList.add("hidden");
}