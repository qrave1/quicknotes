import {Note} from "./notes.tsx";

type Folder = {
    Id: number;
    Name: string;
    Notes: Note[];
};

export default Folder;
