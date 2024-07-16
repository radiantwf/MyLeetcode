import threading


class DiningPhilosophers:
    def __init__(self):
        self.fork_conds = [threading.Condition() for i in range(5)]

    # call the functions directly to execute, for example, eat()
    def wantsToEat(self,
                   philosopher: int,
                   pickLeftFork: 'Callable[[], None]',
                   pickRightFork: 'Callable[[], None]',
                   eat: 'Callable[[], None]',
                   putLeftFork: 'Callable[[], None]',
                   putRightFork: 'Callable[[], None]') -> None:

        with self.fork_conds[philosopher], self.fork_conds[(philosopher + 1) % 5]:
            pickLeftFork()
            pickRightFork()
            eat()
            putLeftFork()
            putRightFork()
