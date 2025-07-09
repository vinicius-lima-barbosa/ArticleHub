const NotFound = () => {
  return (
    <div className="min-h-screen flex items-center justify-center">
      <div className="text-center font-zin-serif">
        <h1 className="text-4xl font-semibold mb-2">
          <span className="text-red-700">404</span> - Not Found
        </h1>
        <p>The page you are looking for does not exist.</p>
        <a href="/" className="text-blue-500 hover:underline">
          Go back to home
        </a>
      </div>
    </div>
  );
};

export default NotFound;
