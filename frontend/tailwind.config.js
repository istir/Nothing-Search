/** @type {import('tailwindcss').Config} */
const config = {
  // darkMode: ["class"],
  content: ["./src/**/*.{html,js,jsx,tsx}"],
  theme: {
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
    extend: {
      colors: {
        border: "hsl(var(--border))",
        input: "hsl(var(--input))",
        ring: "hsl(var(--ring))",
        background: "hsla(var(--background) / var(--tw-bg-opacity))",
        foreground: "hsl(var(--foreground))",
        primary: {
          DEFAULT: "hsla(var(--primary) / var(--tw-bg-opacity))",
          foreground: "hsl(var(--primary-foreground))",
        },
        secondary: {
          DEFAULT: "hsla(var(--secondary) / var(--tw-bg-opacity))",
          foreground: "hsl(var(--secondary-foreground))",
        },
        destructive: {
          DEFAULT: "hsla(var(--destructive) / var(--tw-bg-opacity))",
          foreground: "hsl(var(--destructive-foreground))",
        },
        muted: {
          DEFAULT: "hsla(var(--muted) / var(--tw-bg-opacity))",
          foreground: "hsl(var(--muted-foreground))",
        },
        accent: {
          DEFAULT: "hsla(var(--accent) / var(--tw-bg-opacity))",
          foreground: "hsl(var(--accent-foreground))",
        },
        popover: {
          DEFAULT: "hsla(var(--popover) / var(--tw-bg-opacity))",
          foreground: "hsl(var(--popover-foreground))",
        },
        card: {
          DEFAULT: "hsl(var(--card) / var(--tw-bg-opacity))",
          foreground: "hsl(var(--card-foreground))",
        },
      },
      borderRadius: {
        lg: "var(--radius)",
        md: "calc(var(--radius) - 2px)",
        sm: "calc(var(--radius) - 4px)",
      },
    },
  },
};

export default config;
