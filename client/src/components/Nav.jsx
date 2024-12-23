import { Link } from "react-router-dom";
import logo from '../assets/logo.png';
import BASE_URL from "../main";

function Nav({ name, setName }) {
  const logout = async () => {
    await fetch(BASE_URL+"/api/logout", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });

    setName("");
    location.reload();
  };

  return (
    <nav className="flex items-center justify-between flex-wrap bg-blue-600 p-6">
      <div className="flex items-center flex-shrink-0 text-white mr-2">
        <img
          className="fill-current h-9 w-11 mr-2"
          src={logo}
        />
      </div>
      <div className="w-full block flex-grow lg:flex lg:items-center lg:w-auto">
        <div className="text-sm lg:flex-grow">
          <Link
            to="/"
            className="block mt-4 lg:inline-block lg:mt-0 text-blue-200 hover:text-white mr-4"
          >
            API KEY
          </Link>
          <Link
            to="/rfps"
            className="block mt-4 lg:inline-block lg:mt-0 text-blue-200 hover:text-white mr-4"
          >
            RFPs
          </Link>
          <Link
            to="new_rfp"
            className="block mt-4 lg:inline-block lg:mt-0 text-blue-200 hover:text-white mr-4"
          >
            New RFP
          </Link>
          <Link
            to="/equipment"
            className="block mt-4 lg:inline-block lg:mt-0 text-blue-200 hover:text-white"
          >
            Register Equipment
          </Link>
        </div>
        <div>
          <a
            href="#responsive-header"
            className="block mt-4 lg:inline-block lg:mt-0 text-black mr-4"
          >
            Hello, {name}
          </a>
          <Link
            onClick={logout}
            className="inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-blue-500 hover:bg-white mt-4 lg:mt-0"
          >
            Logout
          </Link>
        </div>
      </div>
    </nav>
  );
}

export default Nav;
