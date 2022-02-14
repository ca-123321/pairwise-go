from ortools.sat.python import cp_model

N = 5 
arrays = [
        [27,2,7,12,22,17],
        [7,8,9,11,1,10],
        [16,28,7,6,24,20],
        [21,4,30,23,7,14],
        [26,5,18,15,29,7],
        [25,31,13,3,19,7],
        [6,3,2,5,1,4],
        [11,21,26,31,16,2],
        [29,19,2,24,14,9],
        [30,2,15,25,10,20],
        [13,23,2,18,8,28],
        [30,29,31,28,27,1],
        [11,6,15,23,19,27],
        [4,9,18,27,25,16],
        [10,27,24,21,13,5],
        [3,27,8,14,20,26],
        [19,17,21,1,18,20],
        [25,28,5,14,17,11],
        [13,30,26,6,9,17],
        [29,16,23,3,17,10],
        [4,24,8,31,17,15],
        [24,26,22,23,1,25],
        [11,22,29,4,13,20],
        [3,9,21,22,15,28],
        [10,22,31,6,18,14],
        [30,16,22,5,8,19],
        [16,12,14,13,1,15],
        [24,12,30,3,18,11],
        [23,12,5,31,20,9],
        [4,26,10,19,28,12],
        [6,21,25,8,29,12],
]


class CellsSolutionPrinter(cp_model.CpSolverSolutionCallback):
    def __init__(self, cells):
        cp_model.CpSolverSolutionCallback.__init__(self)
        self.__cells = cells
        self.__solution_count = 0

    def on_solution_callback(self):
        self.__solution_count += 1
        for row in self.__cells:
            print([self.Value(cell) for cell in row])
        print()

    def solution_count(self):
        return self.__solution_count


def main():
    # Creates the solver
    model = cp_model.CpModel()

    # Create the variables
    cells = []
    for i in range(N * N + N + 1):
        row = []
        for j in range(N + 1):
            row.append(model.NewIntVarFromDomain(cp_model.Domain.FromValues(arrays[i]), 'cell: %i' % (i * (N + 1) + j)))
        cells.append(row)

    # Add the constraints
    # All rows must be different
    for row in cells:
        model.AddAllDifferent(row)

    # All columns must be different
    for i in range(N + 1):
        column = []
        for j in range(N * N + N + 1):
            column.append(cells[j][i])
        model.AddAllDifferent(column)

    # Row order+2 (N+1) must be arranged [1, 2, ... N+1]
    for i, row in enumerate(cells):
        if i == N+1:
            for j in range(N+1):
                model.Add(row[j] == j+1)
        else:
            model.AddAllDifferent(row)

    # Solve the model
    solver = cp_model.CpSolver()
    solution_printer = CellsSolutionPrinter(cells)
    # solver.parameters.enumerate_all_solutions = True
    solver.Solve(model, solution_printer)

    # Statistics
    print('\nStatistics')
    print(f'  conflicts      : {solver.NumConflicts()}')
    print(f'  branches       : {solver.NumBranches()}')
    print(f'  wall time      : {solver.WallTime()} s')
    if solver.parameters.enumerate_all_solutions:
        print(f'  solution count : {solution_printer.solution_count()}')


if __name__ == '__main__':
    main()
