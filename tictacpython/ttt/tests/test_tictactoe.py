import unittest
from tictactoe import Game

class TestTicTacToe(unittest.TestCase):
    def test_create_game(self):
        game = Game()