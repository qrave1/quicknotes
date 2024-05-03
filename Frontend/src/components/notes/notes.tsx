import styles from "./notes.module.css"
import StarterKit from "@tiptap/starter-kit";
import {EditorContent, useEditor} from "@tiptap/react";
import {useEffect, useState} from "react";
import axios, {AxiosResponse} from "axios";
import Folder from "./folders.tsx";
import Folder_modal from "./folder_modal.tsx";
import Note_modal from "./note_modal.tsx";

const authKey = "X-Auth-Token"
const apiBaseUrl = "http://localhost:8080/folders";

export type Note = {
    Id: number;
    Title: string;
    Body: string;
}

let onceFetchNotes = 0;

function Notes() {
    const [folders, setFolders] = useState<Folder[]>([]);
    const [isFolderModalOpen, setIsFolderModalOpen] = useState(false);
    const [isNoteModalOpen, setIsNoteModalOpen] = useState(false);
    const [selectedFolderId, setSelectedFolderId] = useState(-1);
    const [selectedNoteId, setSelectedNoteId] = useState(0);

    const handleShowNotes = (folderId: number) => {
        setSelectedFolderId(folderId);
        fetchNotesOnce();
    };

    const handleSelectNote = (noteId: number) => {
        console.log("noteId", noteId, "----", selectedNoteId)
        for (const f of folders) {
            if (f.Id === selectedFolderId) {
                for (const n of f.Notes) {
                    if (n.Id === selectedNoteId) {
                        console.log("set value = ", n.Id)
                        setSelectedNoteId(n.Id);
                        editor?.commands.setContent(n.Body)
                    }
                }
            }
        }
    }

    const openFolderModal = () => {
        setIsFolderModalOpen(true);
    };

    const closeFolderModal = () => {
        setIsFolderModalOpen(false);
    };

    const openNoteModal = () => {
        setIsNoteModalOpen(true);
    };

    const closeNoteModal = () => {
        setIsNoteModalOpen(false);
    };


    useEffect(() => {
        async function fetchFolders() {
            try {
                const response: AxiosResponse<Folder[]> = await axios.get(apiBaseUrl, {
                    headers: {
                        'Authorization': localStorage.getItem(authKey),
                    }
                });
                if (response.data !== null) {
                    const newFolders: Folder[] = [];
                    for (const f of response.data) {
                        const newFolder: Folder = {
                            Id: f.Id,
                            Name: f.Name,
                            Notes: []
                        }

                        newFolders.push(newFolder);
                    }

                    setFolders(newFolders)
                }
            } catch (error) {
                console.error('Error fetching data: ', error);
            }
        }

        // Call the function
        fetchFolders()
    }, []);


    function fetchNotesOnce() {
        if (onceFetchNotes === 0) {
            onceFetchNotes++

            async function fetchNotes() {
                try {
                    for (let f of folders) {
                        if (f.Id !== null && f.Id !== undefined) {
                            // console.log("making req to", f.Id, "/notes")
                            const response: AxiosResponse<Note[]> = await axios.get(apiBaseUrl + '/' + f.Id + '/notes', {
                                headers: {
                                    'Authorization': localStorage.getItem(authKey),
                                }
                            });
                            // console.log("response /notes", response.data)
                            if (response.data !== null) {
                                for (const n of response.data) {
                                    const newNote: Note = {
                                        Id: n.Id,
                                        Title: n.Title,
                                        Body: n.Body
                                    }

                                    f.Notes.push(newNote)
                                }
                            }
                        }
                    }
                } catch (error) {
                    console.error('Error fetching data: ', error);
                }
            }

            // Call the function
            fetchNotes()
        }
    }


    const handleCreateFolder = (folder_name: string) => {
        const Note: Note[] = [];
        const newFolder: Folder = {
            Id: 0,
            Name: folder_name,
            Notes: Note
        }

        const newFolderRequest = async () => {
            try {
                const response = await axios.post(apiBaseUrl,
                    JSON.stringify({
                        "name": folder_name,
                    }),
                    {
                        headers: {
                            'Authorization': localStorage.getItem(authKey),
                            'content-type': 'application/json'
                        }
                    }
                );
                newFolder.Id = response.data.id
                setFolders(() => ({
                    ...folders,
                    newFolder
                }))
            } catch (error) {
                console.error('Error fetching data: ', error);
            }
        };

        newFolderRequest()
    }

    const handleCreateNote = (folder_id: number, note_name: string) => {
        const newNote: Note = {
            Id: 0,
            Title: note_name,
            Body: ""
        }

        const newNoteRequest = async () => {
            try {
                const response = await axios.post(apiBaseUrl + '/' + folder_id + '/' + 'notes',
                    JSON.stringify({
                        "title": note_name,
                        "body": "",
                    }),
                    {
                        headers: {
                            'Authorization': localStorage.getItem(authKey),
                            'content-type': 'application/json'
                        }
                    }
                );
                newNote.Id = response.data.id
                for (let i = 0; i < folders.length; i++) {
                    if (folders[i].Id === folder_id) {
                        folders[i].Notes.push(newNote)
                    }
                }
            } catch (error) {
                console.error('Error fetching data: ', error);
            }
        };

        newNoteRequest()
    }

    const handleDeleteFolder = (folder_id: number) => {
        const deleteFolderRequest = async () => {
            try {
                const response = await axios.delete(apiBaseUrl + '/' + folder_id,
                    {
                        headers: {
                            'Authorization': localStorage.getItem(authKey),
                            'content-type': 'application/json'
                        }
                    }
                );
                // stub
                if (response.status === 418) {
                    console.log(418)
                }

                setFolders(
                    folders.filter(f => f.Id !== folder_id)
                )
            } catch (error) {
                console.error('Error delete folder: ', error);
            }
        };

        deleteFolderRequest()
    }

    const handleSaveNote = (text: string | undefined, folder: number, note: number) => {
        if (text !== undefined) {
            const saveNote = async () => {
                try {
                    const response = await axios.put(apiBaseUrl + '/' + folder + '/notes/' + note,
                        JSON.stringify({
                            "body": text,
                        }),
                        {
                            headers: {
                                'Authorization': localStorage.getItem(authKey),
                                'content-type': 'application/json'
                            }
                        }
                    );

                    // stub
                    if (response.status === 418) {
                        console.log(418)
                    }
                } catch (error) {
                    console.error('Error save note: ', error);
                }
            };

            saveNote()
        }
    }

    // document.addEventListener('keydown', function (event) {
    //     if (event.key === 's' && (event.ctrlKey || event.metaKey)) {
    //         event.preventDefault();
    //
    //
    //     }
    // });

    useEffect(() => {
        const handleKeyDown = (event: { key: string; metaKey: any; ctrlKey: any; preventDefault: () => void; }) => {
            if (event.key === 's' && (event.metaKey || event.ctrlKey)) {
                event.preventDefault();

                handleSaveNote(editor?.getText(), selectedFolderId, selectedNoteId)
            }
        };

        document.addEventListener('keydown', handleKeyDown);

        return () => {
            document.removeEventListener('keydown', handleKeyDown);
        };
    }, [selectedFolderId, selectedNoteId]);


    const editor = useEditor({
        extensions: [StarterKit],
        content: "",
        editorProps: {
            attributes: {
                class: styles.textEditor,
            }
        }
    });

    return (
        <div className={styles.pageContainer}>
            <Folder_modal isOpen={isFolderModalOpen} onClose={closeFolderModal} onCreateFolder={handleCreateFolder}/>
            <Note_modal isOpen={isNoteModalOpen} onClose={closeNoteModal} folderId={selectedFolderId}
                        onCreateNote={handleCreateNote}/>
            <div className={styles.sidebar}>
                <button className={styles.sidebarButton} onClick={openFolderModal}>Новая папка</button>
                <ul>
                    {folders.map(folder => (
                        <li key={folder.Id}>
                            <div className={styles.folder}>
                                <button onClick={() => handleShowNotes(folder.Id)}>{folder.Name}</button>
                                <button onClick={() => {
                                    setSelectedFolderId(folder.Id);
                                    openNoteModal();
                                }}>+
                                </button>
                                <button onClick={() => handleDeleteFolder(folder.Id)}>X</button>
                            </div>

                            {folder.Id === selectedFolderId ? (
                                <ul>
                                    {folder.Notes.map(note => (
                                        <li key={note.Id + '_' + note.Title} className={note.Id + '_' + note.Title}>
                                            <button onClick={() => {
                                                handleSelectNote(note.Id);
                                                setSelectedNoteId(note.Id)
                                            }}>{note.Title}</button>
                                        </li>
                                    ))}
                                </ul>
                            ) : null
                            }
                        </li>
                    ))}
                </ul>


            </div>
            <div className={styles.editorContainer}>
                <EditorContent editor={editor} className={styles.textEditorContent}/>
            </div>
        </div>

    )
}

export default Notes;
