defmodule Elixirpass do
  @moduledoc """
  Generate a passphrase from the system dictionary.
  """

  @doc """
  Generate a passphrase

  ## Examples

      iex> Elixirpass.generate_passphrase
      "one two three four"

  """
  def generate_passphrase do
    words =
      File.stream!("/usr/share/dict/words")
      |> Stream.map(&String.trim_trailing/1)
      |> Enum.to_list
      |> Enum.shuffle
      |> Enum.take(4)
      |> Enum.join(" ")
    IO.puts words
  end
end
