import React from "react";

import Header from "../components/home/Header";
import HeroHome from "../components/home/HeroHome";
import FeaturesHome from "../components/home/Features";
import FeaturesBlocks from "../components/home/FeaturesBlocks";
import Testimonials from "../components/home/Testimonials";
import Footer from "../components/home/Footer";

function Home() {
  return (
    <div className="flex flex-col min-h-screen overflow-hidden">
      {/*  Site header */}
      <Header />

      {/*  Page content */}
      <main className="flex-grow">
        {/*  Page sections */}
        <HeroHome />
        <FeaturesHome />
        <FeaturesBlocks />
        <Testimonials />
      </main>

      {/*  Site footer */}
      <Footer />
    </div>
  );
}

export default Home;
