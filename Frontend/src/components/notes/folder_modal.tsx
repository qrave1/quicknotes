import React, { useState } from 'react';
import styles from './modal.module.css'

type ModalProps = {
    isOpen: boolean;
    onClose: () => void;
    onCreateFolder: (folderName: string) => void;
};

const Folder_modal: React.FC<ModalProps> = ({ isOpen, onClose, onCreateFolder }) => {
    const [folderName, setFolderName] = useState('');

    const handleCreate = () => {
        onCreateFolder(folderName);
        setFolderName('');
        isOpen = false
    };

    if (!isOpen) return null;

    return (
        <div className={styles.modalOverlay}>
            <div className={styles.modalContent}>
                <h3>Создать папку</h3>
                <input
                    type="text"
                    value={folderName}
                    onChange={(e) => setFolderName(e.target.value)}
                    placeholder="Название"
                />
                <button onClick={handleCreate}>Создать</button>
                <button onClick={onClose}>Закрыть</button>
            </div>
        </div>
    );
};

export default Folder_modal;
