/// <reference types="vite/client" />
interface ImportMetaEnv {
  readonly VITE_API_URL: string;
  // more env variables...
}

// declare module "@tanstack/react-router" {
//   interface Register {
//     router: typeof router;
//   }
// }
