/**
 * 封装的 localStorage 工具模块
 */
const Storage = {
    /**
     * 从 localStorage 获取数据
     * @param key 存储的键
     * @returns 存储的值，或者 null 如果没有找到
     */
    get<T>(key: string): T | null {
        const item = localStorage.getItem(key);
        if (item) {
            try {
                return JSON.parse(item) as T;
            } catch (error) {
                console.error(`Error parsing localStorage item "${key}":`, error);
                return null;
            }
        }
        return null;
    },

    /**
     * 将数据保存到 localStorage
     * @param key 存储的键
     * @param value 要存储的值
     */
    set<T>(key: string, value: T): void {
        try {
            localStorage.setItem(key, JSON.stringify(value));
        } catch (error) {
            console.error(`Error setting localStorage item "${key}":`, error);
        }
    },

    /**
     * 删除 localStorage 中的指定项
     * @param key 要删除的键
     */
    remove(key: string): void {
        localStorage.removeItem(key);
    },

    /**
     * 清空 localStorage
     */
    clear(): void {
        localStorage.clear();
    }
};

export default Storage;