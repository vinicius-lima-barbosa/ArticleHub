import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import SignUpModal from '../register-modal/Index';
import SignInModal from '../login-modal/Index';

const NAV_ITEMS = [
  { name: 'About' },
  { name: 'Sign in' },
  { name: 'Get Started' }
];

const Navigation = () => {
  const navigate = useNavigate();
  const [isSignUpModalOpen, setIsSignUpModalOpen] = useState(false);
  const [isSignInModalOpen, setIsSignInModalOpen] = useState(false);

  const handleNavigation = (itemName: string) => {
    if (itemName === 'Get Started') {
      setIsSignUpModalOpen(true);
    } else if (itemName === 'Sign in') {
      setIsSignInModalOpen(true);
    } else {
      navigate(`/${itemName.toLowerCase().replace(' ', '-')}`);
    }
  };

  return (
    <>
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
                onClick={() => handleNavigation(item.name)}
              >
                {item.name}
              </button>
            ))}
          </div>
        </div>
      </nav>

      <SignInModal
        isOpen={isSignInModalOpen}
        onClose={() => setIsSignInModalOpen(false)}
      />

      <SignUpModal
        isOpen={isSignUpModalOpen}
        onClose={() => setIsSignUpModalOpen(false)}
      />
    </>
  );
};

export default Navigation;
