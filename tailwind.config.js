/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/**/*.templ", "./**/*/go"],
  theme: {
    container: {
      center: true,
      padding: '2rem',
      screens: {
        '2xl': '1400px'
      }
    },
    extend: {
      colors: {
        'brutalist-yellow': '#dea603',
        'brutalist-light-yellow': '#fff3b0',
        'brutalist-black': '#000000',
        'brutalist-grey': '#333333',
        'brutalist-light-grey': '#CCCCCC',
      },
      boxShadow: {
        'brutalist-shadow': '0 4px 6px rgba(0, 0, 0, 0.1), 0 1px 3px rgba(0, 0, 0, 0.08)',
        'brutalist-border': '0 0 0 2px #dea603',
      },
      fontFamily: {
        'brutalist': ['Impact', 'Arial Black', 'sans-serif'],
      },
      animation: {
        'slide-in': 'slideIn 0.7s ease-out',
        'bounce-in': 'bounceIn 0.7s ease-out',
        'slide-in-right': 'slide-in-right 0.7s ease-out',
        'bounce-in-right': 'bounceInRight 1s ease-out forwards',

      },

      keyframes: {
        slideIn: {
          '0%': { transform: 'translateX(-100%)' },
          '100%': { transform: 'translateX(0)' },
        },
        'slide-in-right': {
          '0%': { transform: 'translateX(100%)', opacity: '0' },
          '100%': { transform: 'translateX(0)', opacity: '1' },
        },
        bounceIn: {
          '0%': { transform: 'scale(0.3)', opacity: '0' },
          '50%': { transform: 'scale(1.05)' },
          '100%': { transform: 'scale(1)', opacity: '1' },
        },
        bounceInRight: {
          '0%': { opacity: '0', transform: 'translateX(200px)' },
          '60%': { transform: 'translateX(-30px)' },
          '80%': { transform: 'translateX(10px)' },
          '100%': { opacity: '1', transform: 'translateX(0)' },
        },
      }
    },
    safelist: [],

  }

}
