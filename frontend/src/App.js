import "./App.css";
import { lazy, Suspense } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import LeftNavigation from "./utils/LeftNavigation";
const User = lazy(() => import("./components/user/User"));
const Note = lazy(() => import("./components/note/Note"));
const Home = lazy(() => import("./components/Home"));
const Login = lazy(() => import("./components/user/Login"));
const Register = lazy(() => import("./components/user/Register"));

function App() {
  // const navigate = useNavigate();
  // const goToHome = () => {
  //   navigate("/");
  // };
  return (
    <div className="App">
      {/* <button onClick={goToHome}>Home</button> */}
      <BrowserRouter>
        <LeftNavigation />
        <Suspense fallback={<h1>Loadding Please Wait ... </h1>}>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/user" element={<User />} />
            <Route path="/note" element={<Note />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
          </Routes>
        </Suspense>
      </BrowserRouter>
    </div>
  );
}

export default App;
