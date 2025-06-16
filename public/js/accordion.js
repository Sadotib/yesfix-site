// document.addEventListener('DOMContentLoaded', function () {
//     const accordions = document.querySelectorAll('.accordion-item');

//     accordions.forEach(item => {
//         item.addEventListener('toggle', function () {
//             if (item.open) {
//                 accordions.forEach(other => {
//                     if (other !== item) other.removeAttribute('open');
//                 });
//             }
//         });
//     });
// });

document.addEventListener('DOMContentLoaded', function () {
  const items = document.querySelectorAll('.accordion-item');

  items.forEach(item => {
    const toggle = item.querySelector('.accordion-toggle');
    const content = item.querySelector('.accordion-content');

    toggle.addEventListener('click', () => {
      const isOpen = content.style.maxHeight && content.style.maxHeight !== '0px';

      // Close all others
      items.forEach(otherItem => {
        const otherContent = otherItem.querySelector('.accordion-content');
        if (otherContent !== content) {
          otherContent.style.maxHeight = null;
          otherItem.classList.remove('open');
        }
      });

      // Toggle this one
      if (!isOpen) {
        content.style.maxHeight = content.scrollHeight + "px";
        item.classList.add('open');
      } else {
        content.style.maxHeight = null;
        item.classList.remove('open');
      }
    });
  });
});
