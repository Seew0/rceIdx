import React from "react";
import "tailwindcss/tailwind.css";
import './Home.css'
import rceDummy from './rceDummy.jpg'
import Button from "./Button";
import { FaFile, FaGithub, FaLinkedinIn, FaSpotify} from "react-icons/fa";

function Home() {
    return (
        <div
            className="home min-h-screen flex flex-col justify-center items-center">

            <div className="herobanner">
                <h1 className="font-extrabold text-transparent text-8xl bg-clip-text bg-gradient-to-r from-purple-400 to-pink-600">RCE-IDX</h1>
                <div className="flex flex-wrap justify-center gap-8 ">
                    <Button title="Get Started" link="./codeeditor" />
                </div>

            </div>


            <div className="mt-8 flex justify-center flex-col items-center bannerdisplay">

                <div className="editor1display">
                    <article className="description">
                        <h1 className="text-4xl font-bold mb-8">RCE-IDX: Web App for Efficient Coding</h1>
                        <p className="font-normal text-gray-700 dark:text-gray-400 text-lg">
                            <strong>RCE-IDX</strong> is a web app that allows users to write and run code in multiple programming languages with custom input. With its user-friendly interface, <strong> multiple themes, and font sizes </strong>, along with convenient shortcuts for code execution, Code Editor is a valuable tool for programmers of all levels. The app also includes an <strong> internet status</strong> indicator to ensure users are always connected.
                        </p>


                    </article>
                    <div className="screenshot">
                        <img src={rceDummy} alt="CodePad 1" className="w-1/2 mb-4" />
                    </div>

                </div>

            </div>



            <footer className="bg-gray-800 py-6">
                <div className="container mx-auto px-4">
                    <div className="flex flex-wrap justify-center">
                        <div className="w-full lg:w-6/12 px-4">
                            <h4 className="text-3xl font-semibold text-white">
                                Get in Touch
                            </h4>
                            <div className="mt-6">
                                <a
                                    href="https://github.com/seew0"
                                    className="text-gray-900 bg-white hover:bg-gray-100 border border-gray-200 focus:ring-4 focus:outline-none focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center dark:focus:ring-gray-800 dark:bg-white dark:border-gray-700 dark:text-gray-900 dark:hover:bg-gray-200 mr-2 mb-2"
                                    type="button"
                                    style={{ transition: 'all .15s ease' }}
                                >
                                    <FaGithub className="mr-1" /> Github
                                </a>


                                <a
                                    href="https://open.spotify.com/user/t7ldt174ttnbzlgs3kflptezv?si=1fc96f37e56c4cb0"
                                    className="text-gray-900 bg-white hover:bg-gray-100 border border-gray-200 focus:ring-4 focus:outline-none focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center dark:focus:ring-gray-800 dark:bg-white dark:border-gray-700 dark:text-gray-900 dark:hover:bg-gray-200 mr-2 mb-2"
                                    type="button"
                                    style={{ transition: 'all .15s ease' }}
                                >
                                    <FaSpotify className="mr-1" /> Spotify
                                </a>
                                <a
                                    href="https://linkedin.com/in/devansh-miglani"
                                    className="text-gray-900 bg-white hover:bg-gray-100 border border-gray-200 focus:ring-4 focus:outline-none focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center dark:focus:ring-gray-800 dark:bg-white dark:border-gray-700 dark:text-gray-900 dark:hover:bg-gray-200 mr-2 mb-2"
                                    type="button"
                                    style={{ transition: 'all .15s ease' }}
                                >
                                    <FaLinkedinIn className="mr-1" /> LinkedIn
                                </a>
                                <a
                                    href="https://devanshmiglaniresume.tiiny.site/"
                                    className="text-gray-900 bg-white hover:bg-gray-100 border border-gray-200 focus:ring-4 focus:outline-none focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center dark:focus:ring-gray-800 dark:bg-white dark:border-gray-700 dark:text-gray-900 dark:hover:bg-gray-200 mr-2 mb-2"
                                    type="button"
                                    style={{ transition: 'all .15s ease' }}
                                >
                                    <FaFile className="mr-1" /> Resume
                                </a>
                            </div>
                        </div>
                    </div>
                    <hr className="my-6 border-gray-700" />
                    <div className="flex flex-wrap items-center md:justify-between justify-center">
                        <div className="w-full  px-4">
                            <div className="text-sm text-white font-semibold py-1">
                                © {new Date().getFullYear()} seew0
                            </div>
                        </div>

                    </div>
                </div>
            </footer>
        </div >



    );
}

export default Home;
