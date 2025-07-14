import { useNavigate } from 'react-router-dom';

const NAV_ITEMS = [
  { name: 'About' },
  { name: 'Sign in' },
  { name: 'Get Started' }
];

const Navigation = () => {
  const navigate = useNavigate();

  return (
    <nav className="bg-beige fixed top-0 z-30 w-full border-b border-gray-900 shadow-md h-20">
      <div className="px-6 py-5 flex items-center justify-between max-w-4xl mx-auto">
        <h1
          onClick={() => navigate('/')}
          className="font-zin-serif text-3xl font-extrabold cursor-pointer"
        >
          Article Hub
        </h1>
        <div className="font-helvetica-now-var hidden md:flex items-center space-x-4">
          {NAV_ITEMS.map((item) => (
            <button
              key={item.name}
              className={`text-sm cursor-pointer ${item.name === 'Get Started' ? 'bg-gray-900 text-white px-4 py-2 rounded-full' : ''}`}
              onClick={() =>
                navigate(`/${item.name.toLowerCase().replace(' ', '-')}`)
              }
            >
              {item.name}
            </button>
          ))}
        </div>
      </div>
    </nav>
  );
};

export default Navigation;
