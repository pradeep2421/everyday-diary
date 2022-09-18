import { useNavigate } from "react-router-dom";

const Home = () => {
  const navigate = useNavigate();
  const goToLoginPage = () => {
    navigate("/login");
  };
  const goToRegisterPage = () => {
    navigate("/register");
  };
  return (
    <div>
      <h1 className="home">Home</h1>
      <button className="button" onClick={goToLoginPage}>
        Login
      </button>
      <button className="button" onClick={goToRegisterPage}>
        Register{" "}
      </button>
    </div>
  );
};

export default Home;
