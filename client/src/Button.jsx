import { Link } from 'react-router-dom';

function Button({ title, link }) {
    return (
        <Link to={link} className="cardhome bg-white shadow-lg overflow-hidden p-2">
            <div className="px-1 py-1">
                <div className=" font-bold text-xl mb-1">{title}</div>
            </div>
        </Link >
    );
}

export default Button;
