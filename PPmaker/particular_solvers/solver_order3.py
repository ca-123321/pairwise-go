from ortools.sat.python import cp_model

N = 3 
arrays = [
    [2, 5, 8, 11],
    [1, 5, 6, 7],
    [4, 5, 10, 12],
    [3, 5, 9, 13],
    [1, 2, 3, 4],
    [2, 7, 10, 13],
    [2, 6, 9, 12],
    [1, 11, 12, 13],
    [4, 7, 9, 11],
    [3, 6, 10, 11],
    [1, 8, 9, 10],
    [3, 7, 8, 12],
    [4, 6, 8, 13]
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
