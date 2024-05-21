import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import CharacterApp from "./CharacterApp";
import PeopleApp from "./PeopleApp";
import GamePageApp from "./GamePageApp";
import CharacterPageApp from "./CharacterPageApp";
import PeoplePageApp from "./CharacterPageApp";
import './index.css';
import Login from "./components/LoginPage"
import Register from "./components/Register"
import reportWebVitals from "./reportWebVitals";
import { ClerkProvider, RedirectToSignIn, SignIn, SignUp, SignedIn, SignedOut } from "@clerk/clerk-react";
import { BrowserRouter, Routes, Route, useNavigate } from "react-router-dom";
import ProtectedPage from "./ProtectedPage";


const clerkPubKey = process.env.VITE_CLERK_PUBLISHABLE_KEY;

const root = ReactDOM.createRoot(document.getElementById("root"));

const ClerkWithRoutes = () => {
  const navigate = useNavigate();

  return (
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/characters" element={<CharacterApp />} />
        <Route path="/people" element={<PeopleApp />} />
        <Route path="/:id" element={<GamePageApp />} />
        <Route path="/characters/:id" element={<CharacterPageApp />} />
          <Route path="/people/:id" element={<PeoplePageApp />} />
        <Route
          path="/user/sign_in"
          element={<Register />}
        />
        <Route
          path="/user/login"
          element={<Login />}
        />
        <Route
          path="/protected"
          element={
            <>
            <SignedIn>
              <ProtectedPage />
            </SignedIn>
            <SignedOut>
              <RedirectToSignIn />
            </SignedOut>
            </>
          }
        />
      </Routes>
  );
};

root.render(
  <React.StrictMode>
    <BrowserRouter>
      <ClerkWithRoutes />
    </BrowserRouter>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
