import { NavLink } from "react-router-dom";

const LeftNavigation = () => {
  return (
    <div className="header">
      <nav className="nav-menu1 active1">
        <div>
          <NavLink className="navlink-class" to="/">
            <div className="left-nav-text" id="left-nav-id">
              <i className="fa fa-user"></i>
              <span>HOME</span>
            </div>
          </NavLink>
        </div>
      </nav>
    </div>
  );
};

export default LeftNavigation;
