const HeroSection = () => {
  return (
    <section className="bg-beige flex-1 bg-cover bg-center flex items-center justify-center">
      <div className="px-4 flex flex-col-reverse md:flex-row items-center">
        <div className="text-start mt-8 md:mt-8">
          <h1 className="font-zin-serif text-8xl mt-20">Welcome to...</h1>
          <h1 className="font-zin-serif text-8xl italic ml-6">Article Hub</h1>
          <p className="font-helvetica-now-var mt-4 text-2xl">
            Your go-to platform for the latest articles and insights.
          </p>
        </div>
      </div>
    </section>
  );
};

export default HeroSection;
