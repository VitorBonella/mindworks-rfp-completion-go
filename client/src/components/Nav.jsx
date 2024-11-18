import { Link } from "react-router-dom";

function Nav({ name, setName }) {
  const logout = async () => {
    await fetch("http://localhost:8000/api/logout", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });

    setName("");
  };

  return (
    <nav className="flex items-center justify-between flex-wrap bg-blue-600 p-6">
      <div className="flex items-center flex-shrink-0 text-white mr-6">
        <img
          className="fill-current h-8 w-8 mr-2"
          src="https://cdn-icons-png.flaticon.com/512/5208/5208460.png"
        />
      </div>
      <div className="w-full block flex-grow lg:flex lg:items-center lg:w-auto">
        <div className="text-sm lg:flex-grow">
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
