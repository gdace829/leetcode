#include <sys/ipc.h>
#include <sys/shm.h>
#include <iostream>

int main(int argc, char const *argv[])
{
    key_t key = ftok("path_to_file", 0);                 // 获取唯一键值
    int shm_id = shmget(key, 10, IPC_CREAT | 0600);      // 创建或获取共享内存段
    void *shared_memory_ptr = shmat(shm_id, nullptr, 0); // 映射共享内存
    if (shared_memory_ptr == reinterpret_cast<void *>(-1))
    {
        std::cout << "err" << std::endl;
    }
    std::int32_t *data = static_cast<std::int32_t *>(shared_memory_ptr);
    *data = 42; // 写入数据
    std::int32_t *data1 = static_cast<std::int32_t *>(shared_memory_ptr);
    std::cout << *data1 << std::endl;
    shmdt(shared_memory_ptr);    // 解除映射
    shmctl(shm_id, IPC_RMID, 0); // 删除共享内存段
    return 0;
}
