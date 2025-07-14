import Footer from '../components/landing-page/footer/Index';
import HeroSection from '../components/landing-page/hero-section/Index';
import Navigation from '../components/landing-page/navigation/Index';

const LandingPage = () => {
  return (
    <div className="flex flex-col h-screen overflow-hidden">
      <Navigation />
      <HeroSection />
      <Footer />
    </div>
  );
};

export default LandingPage;
