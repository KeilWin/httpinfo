import { Router, Route } from "@solidjs/router";

import Header from "./pages/Header";
import Footer from "./pages/Footer";
import HomePage from "./pages/home/HomePage";

import "./styles/app.css";

export default function App() {
  return (
    <div class="app">
        <Header />
        <Router>
          <Route path="/" component={HomePage}/>
        </Router>
        <Footer />
    </div>
  );
}
