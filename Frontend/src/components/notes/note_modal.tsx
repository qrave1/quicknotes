import React, {useState} from 'react';
import styles from './modal.module.css'

type ModalProps = {
    isOpen: boolean;
    onClose: () => void;
    folderId: number;
    onCreateNote: (folderId: number, noteName: string) => void;
};

const Note_modal: React.FC<ModalProps> = ({isOpen, onClose, folderId, onCreateNote}) => {
    const [noteName, setNoteName] = useState('');

    const handleCreate = () => {
        onCreateNote(folderId, noteName);
        setNoteName('');
        isOpen = false
    };

    if (!isOpen) return null;

    return (
        <div className={styles.modalOverlay}>
            <div className={styles.modalContent}>
                <h3>Создать запись</h3>
                <input
                    type="text"
                    value={noteName}
                    onChange={(e) => setNoteName(e.target.value)}
                    placeholder="Название"
                />
                <button onClick={handleCreate}>Создать</button>
                <button onClick={onClose}>Закрыть</button>
            </div>
        </div>
    );
};

export default Note_modal;
